## Running the App Locally

### Step 1: Start Local Supabase

From the project root, start the local Supabase instance:

```bash
cd backend/internal
bun run supabase start
```

Once started, the CLI will output connection details including URLs and keys. Keep this output handy for the next step.

---

### Step 2: Configure Environment

Create `backend/.local.env` using `backend/.env.template` as a reference:

```
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=host.docker.internal
DB_PORT=54322
DB_NAME=postgres
PORT=8080
SUPABASE_URL=http://host.docker.internal:54321
SUPABASE_ANON_KEY=sb_publishable_ACJWlzQHlZjBrEguHvfOxg_3BJgxAaH
SUPABASE_SERVICE_ROLE_KEY=sb_secret_N7UND0UgjKTVK-Uodkm0Hg_xSvEMPvz
DB_SSLMODE=disable
```

Replace the placeholder values with the keys from your `supabase start` output.

> **Important:** The Supabase output may show `127.0.0.1` for URLs and hosts. Replace any instance of `127.0.0.1` with `host.docker.internal` — otherwise Docker containers will reference their own localhost rather than your machine.

---

### Step 3: Set Environment Mode

In the root `.env` file, ensure the `ENVIRONMENT` variable is set:

```
ENVIRONMENT=development
```

This tells the backend to load `backend/.local.env`. Setting it to `production` would load `backend/.env` instead.

---

### Step 4: Start the Application

From the project root:

```bash
make up
```

This spins up both the backend and frontend.

---

### Step 5: Verify It's Working

1. Open Supabase Studio at `http://127.0.0.1:54323` to view your local database
2. Create a test entity via the API
3. Confirm the data appears in Supabase Studio

---

### Stopping the Application

To stop the Docker containers:

```bash
make down
```

To stop the local Supabase instance:

```bash
cd backend/internal
bun run supabase stop
```

--- 