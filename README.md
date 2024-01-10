
## Obtaining credentials.json
To use Google APIs in your Go application, you need to create a project in the Google Developers Console and set up OAuth 2.0 credentials. Follow these steps to get started:

1. Go to the Google Developers Console
   Action: Visit the [Google Developers Console](https://console.cloud.google.com/welcome).
2. Create a Project
   Action: Click on "Select a project" at the top of the page, then click on "NEW PROJECT" and follow the prompts to create a new project.
3. Enable the Google Apps Script API
   Action: In the dashboard, navigate to "Library" in the left-hand menu.
   Search: Look for "Google Apps Script API" and enable it.
4. Configure the OAuth Consent Screen
   Action: Go to "OAuth consent screen" on the left sidebar.
   User Type: Select the User Type (usually "External") and click "Create".
   Details: Fill in the required fields (like app name, user support email, etc.).
5. Create OAuth 2.0 Credentials
   Action: Navigate to "Credentials" on the left sidebar and click on "Create Credentials" at the top of the page.
   Credential Type: Choose "OAuth client ID".
   Application Type: Select "Web application" or "Other" depending on your use case.
   Redirect URIs: Add the necessary "Authorized redirect URIs" if required (for web applications).
   Create: Click "Create".
6. Download the Credentials
   Action: Once the credentials are created, a confirmation screen will appear with your client ID and client secret.
   Download: Click on "Download JSON" on the right side of your OAuth 2.0 Client IDs. This will download the credentials.json file.
7. Use the Credentials in Your Application
   Placement: Place the credentials.json file in the same directory as your Go application or provide the path to this file in your code.
   Note: After completing these steps, your Go application will be able to use these credentials to authenticate with Google's services, including the Apps Script API. Handle these credentials securely and do not expose them publicly