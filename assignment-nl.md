# Toetsopdracht DevOps: NSE Thermometer
_maart 2026_

De Nationale Studenten Enquête is een grootschalig landelijk tevredenheidsonderzoek waarin studenten hun mening geven over hun HBO-opleiding en onderwijsinstelling. Waarschijnlijk heb jij deze enquête ook wel eens ingevuld.

HBO-ICT gebruikt verschillende systemen om de NSE te promoten. Een van deze systemen is de NSE Stretch Goal Thermometer. De applicatie die op verschillende schermen binnen Saxion zichtbaar is en het huidige aantal stemmen en de behaalde stretch goals weergeeft.

In deze examenopdracht is het jouw taak om deze applicatie goed draaiende te krijgen in de cloud. Alles moet zo geautomatiseerd worden zodat je applicatie eenvoudig uitgerold kan worden vanuit Gitlab naar een schaalbare cloud gebaseerde oplossing.

### Voorwaardelijke Eisen

-   Dit vak wordt afgesloten met een assessment. Je kunt enkel deelnemen aan het assessment wanneer al je huiswerkopdrachten afgetekend zijn. Uiteraard kun je ondertussen alvast wel een beginnetje maken met deze examenopdracht. De herkansingsopdracht is een andere opdracht, dus het heeft geen zin om aan deze eindopdracht te werken als je je huiswerk niet op tijd gaat afkrijgen.

-   Voor elke onderdeel van de opdracht (zie hieronder) leg je uit hoe je het probleem hebt opgelost en waarom je het op deze manier hebt opgelost. Dit beschrijf je in je eigen woorden worden op een manier die te begrijpen is voor de docent die je werk beoordeelt. Je werkt je documentatie uit in een Markdown-bestand in je repository. Je kunt ook gebruik maken van commentaar in de source code (in dit geval verwijs je in het Markdown-bestand naar de locaties van deze comments).

-   Je documentatie is **maximaal** 1800 woorden (2 A4’tjes).

-   Zorg ervoor dat je een repository gebruikt in de Saxion Gitlab organisatie, gemaakt via [https://repo.hboictlab.nl](https://repo.hboictlab.nl/).

-   Upload een gezipte export van je Gitlab-project naar Brightspace (Settings → General → Advanced → Export Project). 

    **Vermeld ook de URL van je repository in het opmerkingenveld op Brightspace.**

### Toetsregels

-   De deadline voor het inleveren van de opdracht is **maandag 6 april 23:59 (begin week 3.9)**.

-   De opdracht wordt gemaakt in groepjes van twee studenten die zich hebben aangemeld bij de docent. Groepjes worden eind week 6 gevormd, waarbij groepsleden allebei evenveel oefenopdrachten afgetekend moeten hebben om te mogen samenwerken (maximaal 1 week verschil).

-   Je cijfer wordt bepaald op basis van het assessment dat plaatsvindt nadat je je opdracht hebt ingeleverd. Zorg dus dat je alle onderdelen die je gemaakt hebt goed begrijpt, er zullen ook theoretische vragen worden gesteld over je gemaakte oplossing.

-   Je krijgt alleen punten voor onderdelen van je werk die je tijdens het assessment kunt uitleggen. Bijvoorbeeld: als je code werkt, maar je kunt niet uitleggen wat het doet of waarom het zo werkt, dan telt dat onderdeel van de opdracht niet mee voor je eindcijfer.

-   Het is **niet** toegestaan om hulp te krijgen van iemand buiten je toetsgroepje, behalve van je docent. Het is ook **niet** toegestaan om hulp te verlenen aan iemand anders.

-   Je mag online posts, artikelen, tutorials, boeken, video\'s gebruiken. Je moet wel [verwijzingen toevoegen](https://saxionbibliotheek.libguides.com/c.php?g=707376&p=5098767) voor alle bronnen en codefragmenten die je in je tekst/code gebruikt.

-   Je mag Generatieve AI transparant en verantwoord gebruiken volgens de [handreiking AI-gebruik voor studenten van Saxion](https://saxion.sharepoint.com/sites/o365-kpf-000001/Gedeelde%20%20documenten/Forms/AllItems.aspx?id=%2Fsites%2Fo365%2Dkpf%2D000001%2FGedeelde%20%20documenten%2FGenAI%20student%20guide%2Epdf&parent=%2Fsites%2Fo365%2Dkpf%2D000001%2FGedeelde%20%20documenten). Je werk bevat een verantwoording over het gebruik van AI in je proces volgens de APA-richtlijnen. Je blijft zelf altijd verantwoordelijk voor de inhoud van je werk. Je moet duidelijk aangeven welke prompts je hebt gebruikt en welke delen van het gegenereerde antwoord je daadwerkelijk hebt overgenomen of gebruikt.

## Examenopdracht

### Deel 1: Development en kwaliteit

#### 1. Inrichten Git proces
Tijdens het uitvoeren van de opdracht verwachten we van je dat je een gedegen softwareontwikkelingsproces volgt, zoals behandeld in week 1. Dit omvat dingen zoals het bijhouden van taken met behulp van Gitlab issues, het gebruik van Git om samen te werken aan de code en het gebruik van branches en merge request om code reviews te kunnen doen. Dit zou allemaal duidelijk moeten blijken uit je Gitlab-repository.

Aanvullend kun je overwegen om je werkproces verder te professionaliseren door gebruik te maken van:

- Rebasing om onnodige commits samen te voegen (commit squashing).
- Git hooks (voor security checks of andere onderdelen die je lokaal wilt controleren. Denk bijvoorbeeld aan [GitLeaks](https://github.com/gitleaks/gitleaks)).
- Gestandaardiseerde commit messages, bijvoorbeeld door gebruik van [Conventional commits](https://www.conventionalcommits.org/en/v1.0.0/).

#### 2. Bestuderen applicaties en lokaal werkend krijgen
In de repository vind je een map met daarin een backend applicatie. Dit is een Go-applicatie die door ons wordt geleverd en die gegevens in JSON-formaat teruggeeft. In de `README.md` in de backend map vind je meer informatie over het project. Installeer de dependencies en kijk of je de applicatie lokaal werkend krijgt. Voer ook de andere commando's van de `README.md` uit voor testen, linting, security checking, etc.

Om te controleren of de applicatie correct werkt, kun je de endpoints van de API aanroepen met behulp van Postman (of een andere tool naar keuze). Het bestand `endpoints.postman_collection.json` kun je importeren in Postman en bevat alle endpoints van de API.

**Tip:** De huidige implementatie van de backend heeft een security vulnerability (SQL injectie). Los dit probleem op, maar bewaar ook de oude code. We willen straks tijdens de pipeline fase ook kunnen testen dat de pipeline faalt bij een security vulnerability.

In de frontend map vind je onze frontend applicatie, gebouwd met Angular. Bestudeer de `README.md` en kijk of je de applicatie kunt installeren en draaien. Kijk ook of je de verschillende `npm run` commandos kunt uitvoeren.

Neem ook de tijd om de code van de backend en de frontend goed te bestuderen. Begrijp welke talen gebruikt worden. Dit helpt je om later de goede keuzes bij docker, CI/CD en deployment te maken. Bijvoorbeeld: 
- heb je een development omgeving nodig voor alle acties?
- wordt de taal gecompileerd of geïnterpreteerd?
- hoe voeg je environment variabelen toe?
- etc.

#### 3. Seed script
Als je de backend applicatie uitvoert en naar [http://localhost:8000/api/goals](http://localhost:8000/api/goals) gaat, krijg je een lijst met stretch goals te zien die op dit moment in de database aanwezig zijn. Initieel is de database leeg.

De map `seed-application` bevat data bestanden die moeten worden ingelezen en toegevoegd aan de backend door middel van een POST request. Het data formaat wordt uitgelegd in de `README.md` in de `seed-application` map.

Update het Bash script `add_data.sh` in de `seed-application` map, waarmee je de data inleest en met behulp van `curl` de stretch goals toevoegt aan de backend (ook wel 'seeding' genoemd in het Engels).

**Tips:**
- Zorg ervoor dat je het content type instelt op `application/json` in je Curl commando.
- Maak je script dynamisch zodat bestanden met andere namen eenvoudig ingelezen kunnen worden, door de bestandsnaam mee te geven als parameter aan je script.
- Zorg dat je script configureerbaar is door gebruik te maken van environment variabelen.
- Zorg voor nette foutafhandeling.

#### 4. Containerisatie van de applicaties

De volgende stap is om de hele applicatie te containeriseren met Docker en Docker Compose. Het doel is dat alle onderdelen (frontend, backend, database en seed script) draaien in containers, zodat je op je eigen machine geen enkele taal, runtime of framework hoeft te installeren.

In de `README.md` bestanden van de frontend en backend staat beschreven welke configuratiemogelijkheden er zijn.

We willen hierbij:
- Losse containers voor backend, frontend, database en seed script.
- Microsoft SQL Server als database gebruiken. De 2022 versie volstaat voor je oplossing. De `README.md` van de backend legt uit welke variabelen je nodig hebt om de databasekoppeling aan te passen.
- Nginx of Apache HTTP Server gebruiken als webserver voor je frontend. Je hoeft **geen** reverse proxy, load balancer of HTTPS toe te voegen.
- Dat het seed script uitgevoerd wordt nadat de backend en de database gestart zijn.
- De images zo klein mogelijk houden. Kies dus handige base images en kijk naar andere mogelijkheden om je images klein te houden.
- De configuratie zo flexibel mogelijk houden, zodat we delen van de docker compositie later opnieuw kunnen gebruiken bij het uitrollen naar de cloud. Denk hierbij aan environment variabelen, .env bestanden, dockerignore, multi-stage builds, docker secrets, etc.

Het is de bedoeling om Docker al het werk te laten doen, maar mocht je toch meerdere stappen lokaal moeten uitvoeren voordat je de compositie kunt starten (`docker compose up`), werk deze stappen dan uit in een shell script en zet het script in je repository.

#### 5. Continuous Integration

Implementeer een ontwikkelstraat met Continuous Integration voor de gehele applicatie. We willen dat het CI gedeelte draait na elke **push** of **merge request** op de **main branch**. Je CI pipeline zou uit een aantal stappen moeten bestaan:

- Backend:
    - Dependencies installeren, test, lint, security check en code coverage.
    - De code coverage website moet als GitLab pages worden aangeboden.
- Frontend:
    - Dependencies installeren, frontend builden en testen.
    - Build moet te downloaden zijn als artefact in Gitlab.
    - Test rapport moet zichtbaar zijn via Gitlab Pages

Probeer de ontwikkelstraat zo efficiënt mogelijk te maken door jobs in parallel te draaien en door caching te gebruiken. Test ook of de security check goed werkt (zie stap 1 van de examenopdracht).

### Deel 2: Deployment
Nu gaan we de applicatie beschikbaar maken met behulp van AWS. We raden je aan om dit stapsgewijs op te bouwen en eenvoudig te beginnen (dus nog zonder load balancers, auto scaling, etc.). In stap 6 en 7 ga je deze eenvoudige architectuur bouwen met Terraform en Gitlab CI/CD. Vervolgens ga je in stap 8 je architectuur uitbouwen om het *highly available* te maken.

Zorg ervoor dat je een diagram maakt van je architectuur in AWS en deze opneemt in je documentatie. Maak hiervoor gebruik van [draw.io](https://draw.io/) en gebruik de AWS 2026 shapes (klik op *+ More Shapes > AWS 2026*).

#### 6. Maak een infrastructuur om de applicatie op te draaien
Creëer een geschikte infrastructuur voor de applicatie in AWS met behulp van Terraform. De infrastructuur hoeft (nog) niet *highly available* en schaalbaar te zijn. Als dit niet lukt met Terraform mag je handmatig de infrastructuur aanmaken (hiermee verlies je natuurlijk wel punten). Je Terraform-configuratie moet in de meegeleverde `infra` map worden geplaatst.

Zorg er in deze stap voor dat de applicatie draait op de AWS architectuur. De applicatie hoeft nog niet uitgerolt te worden met Terraform naar de zojuist gemaakte infrastructuur. Dit ga je doen in stap 7 met Continuous Deployment. 

#### 7. Automatiseer de deployment van de applicaties
Automatiseer de deployment van de backend zodat, wanneer we in GitLab een nieuwe tag toevoegen aan de main branch, de Gitlab CI:

-   De benodigde Docker images bouwt voor de backend
-   De image(s) upload naar de [private container registry](https://docs.gitlab.com/ee/user/packages/container_registry/index.html) van je Gitlab repository
-   De images gebruikt in AWS, zodat de nieuwe versie van de backend in de cloud draait
-   De gebouwde frontend web files uploadt als een statische website (bijvoorbeeld in S3 of in een EC2 instantie)

Er zijn verschillende manieren om de backend te deployen naar je cloudomgeving, bijvoorbeeld via SSH. Zorg ervoor dat je secrets (zoals bijvoorbeeld SSH private keys) veilig opgeslagen zijn.

**Tip:**

- Indien je gebruik maakt van SSH om geautomatiseerd de backend te deployen naar een server, controleert de client of de server een bekende server is. Om dit te automatiseren kun je de **public key** van de server opslaan als omgevingsvariabele in je CI/CD-omgeving en deze toevoegen aan het `known_hosts` bestand:

    ```bash
    echo "$SSH_KNOWN_HOST" >> ~/.ssh/known_hosts
    ```

- Haal de public key eenmalig op door dit lokaal uit te voeren op de server en de uitvoer op te slaan als variabele:

    ```bash
    ssh-keyscan -H <jouw-server-ip>
    ```

**Aanvullende ideeën:**
- Je kunt manual steps opnemen in je pipeline om de infrastructuur aan te maken of weer neer te halen (tear down).

### 8. Maak de backend highly available

Breid je Terraform-configuratie uit met een high availability en scalable backend. Als je frontend op EC2 draait, breid je de schaalbaarheid ook daar op toe. Vergeet niet om de frontend aan te passen om verbinding te maken met je nieuwe infrastructuur. Om te valideren of je infrastructuur correct is ingesteld, check je of je frontend de stretch goals laat zien.

Eisen:
- Het aantal backend machines welke gebruikt worden schaalt automatisch bij meer load.
- Er zijn machines in meerdere availability zones die de requests afhandelen.

Vergeet niet je architectuurdiagram bij te werken, zodat de uiteindelijke architectuur opgenomen wordt in je documentatie.
