# Docker Installation on Linux

*Part of VPS setup series - Last updated: August 2025*

Quick guide for installing Docker on a Linux VPS. This follows the official Docker installation method.

## Installation Steps

1. **Download the Docker installation script**
   ```bash
   curl -fsSL https://get.docker.com -o get-docker.sh
   ```

2. **Run the installation script**
   ```bash
   sudo sh get-docker.sh
   ```

3. **Add your user to the Docker group**
   ```bash
   sudo usermod -aG docker $USER
   ```

4. **Log out and back in** (or run `newgrp docker`) to apply the group changes

5. **Verify the installation**
   ```bash
   docker --version
   docker run hello-world
   ```

## What This Does

- **Official script**: Uses Docker's official installation script for your Linux distribution
- **User permissions**: Adding your user to the `docker` group lets you run Docker commands without `sudo`
- **Group refresh**: You need to refresh your group membership for the permissions to take effect

## Next Steps

After Docker is installed, you might want to:

- Set up Docker Compose for multi-container applications
- Configure Docker to start on boot: `sudo systemctl enable docker`
- Learn about container management and best practices

## Troubleshooting

If you get permission errors after adding your user to the docker group, try:
```bash
newgrp docker
```

Or simply log out and log back in to refresh your group membership.

---

*This is part of my VPS setup learning series. See the [main VPS guide](./1-setup.md) for initial server configuration.*