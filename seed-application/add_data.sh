#!/usr/bin/env bash

#db cleaning
# sqlite3 app.db "DELETE FROM goals; DELETE FROM percentages;"

BACKEND_URL="${BACKEND_URL:-http://localhost:8000}"
URL="${BACKEND_URL}/api/goals"

MMD="goals.mmd"
MMV="goals.mmv"

declare -A values
while IFS= read -r line || [ -n "$line" ]; do
  if [ -n "${line// }" ]; then
    read -r key rest <<< "$line"
    if [ -n "$key" ]; then
      values["$key"]="$rest"
    fi
  fi
done < "$MMV"

declare -A target_ids
declare -A goal_ids
while IFS= read -r line || [ -n "$line" ]; do
  if [ -n "${line// }" ]; then
    read -r w1 eid w3 prop w5 vid <<< "$line"
    if [ "$w1" = entry ] && [ "$w3" = property ] && [ "$w5" = value_id ]; then
      if [ "$prop" = target ]; then
        target_ids["$eid"]="$vid"
      elif [ "$prop" = goal ]; then
        goal_ids["$eid"]="$vid"
      fi
    fi
  fi
done < "$MMD"

for eid in "${!target_ids[@]}"; do
  tv="${target_ids[$eid]}"
  gv="${goal_ids[$eid]:-}"
  if [ -n "$gv" ] && [ -n "${values[$tv]:-}" ] && [ -n "${values[$gv]:-}" ]; then
    target_text="${values[$tv]}"
    goal_name="${values[$gv]}"
    if [[ "$target_text" =~ ([0-9]+) ]]; then
      pct="${BASH_REMATCH[1]}"
      n="${goal_name//\'/\'\'}"
      n="${n//\\/\\\\}"
      n="${n//\"/\\\"}"
      json="{\"name\":\"${n}\",\"targetPercentage\":${pct}}"
      curl -sS -X POST -H "Content-Type: application/json" -d "$json" "$URL"
      echo
    fi
  fi
done
