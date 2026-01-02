# Deploying Habit Tracker to VPS with Nginx

## Prerequisites
- VPS with Nginx already running
- Domain `leoclaudio.dev` already configured
- SSH access to your VPS

## Step 0: Configure Google Analytics (Optional)

1. Go to [Google Analytics](https://analytics.google.com/)
2. Create a new property for `habit.leoclaudio.dev`
3. Get your Measurement ID (format: `G-XXXXXXXXXX`)
4. Create a `.env` file in your project root:
   ```bash
   cp .env.example .env
   ```
5. Replace `G-XXXXXXXXXX` with your actual Measurement ID in `.env`:
   ```
   VITE_GA_MEASUREMENT_ID=G-YOUR-ACTUAL-ID
   ```

## Step 1: Build Your Application

```bash
npm run build
```

## Step 2: Prepare VPS Directory

```bash
# Create directory for habit tracker
sudo mkdir -p /var/www/habit-tracker

# Upload built files from local machine
scp -r dist/* user@your-vps-ip:/var/www/habit-tracker/

# Set proper permissions
sudo chown -R www-data:www-data /var/www/habit-tracker
sudo chmod -R 755 /var/www/habit-tracker
```

## Step 3: Create Nginx Configuration

```bash
sudo nano /etc/nginx/sites-available/habit.leoclaudio.dev
```

Add this configuration:

```nginx
server {
    listen 80;
    server_name habit.leoclaudio.dev;
    root /var/www/habit-tracker;
    index index.html;

    # Handle React Router - serve index.html for all routes
    location / {
        try_files $uri $uri/ /index.html;
    }

    # Cache static assets
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg)$ {
        expires 1y;
        add_header Cache-Control "public, no-transform";
    }

    # Security headers
    add_header X-Frame-Options "DENY" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header Referrer-Policy "strict-origin-when-cross-origin" always;

    # Gzip compression
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_types text/plain text/css text/xml text/javascript application/javascript application/xml+rss application/json;
}
```

## Step 4: Enable the Site

```bash
# Create symbolic link to enable site
sudo ln -s /etc/nginx/sites-available/habit.leoclaudio.dev /etc/nginx/sites-enabled/

# Test Nginx configuration
sudo nginx -t

# Reload Nginx
sudo systemctl reload nginx
```

## Step 5: Configure DNS

Add an A record to your DNS provider:
- **Type**: A
- **Name**: habit
- **Value**: your-vps-ip-address
- **TTL**: 300 (or default)

## Step 6: Add SSL Certificate

```bash
sudo certbot --nginx -d habit.leoclaudio.dev
```

## Step 7: Verify Deployment

1. Wait for DNS propagation (1-5 minutes)
2. Visit `https://habit.leoclaudio.dev`
3. Your habit tracker should be live!

## Final Setup

Your server now hosts:
- `https://leoclaudio.dev` → Your personal site
- `https://habit.leoclaudio.dev` → Your habit tracker

## Optional: Deployment Script

Create a simple deployment script for future updates:

```bash
#!/bin/bash
# deploy-habit.sh
echo "Building application..."
npm run build

echo "Uploading to server..."
scp -r dist/* user@your-vps-ip:/var/www/habit-tracker/

echo "Setting permissions..."
ssh user@your-vps-ip "sudo chown -R www-data:www-data /var/www/habit-tracker"

echo "Deployment complete!"
```

Make it executable:
```bash
chmod +x deploy-habit.sh
```
