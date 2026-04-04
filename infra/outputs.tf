output "alb_dns" {
  description = "ALB DNS"
  value       = aws_lb.main.dns_name
}

output "frontend_url" {
  description = "S3 static website URL"
  value       = aws_s3_bucket_website_configuration.frontend.website_endpoint
}

output "db_private_ip" {
  description = "Private IP of the MSSQL EC2 instance"
  value       = aws_instance.db.private_ip
}
