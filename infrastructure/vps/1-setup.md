# VPS Setup Guide - Learning from Scratch

*Last updated: August 2025*

Got my hands on a VPS recently and wanted to document what I learned while setting it up. This guide covers the essential security and setup steps I wish I had known from the start.

---

## Starting Fresh

First things first - update everything. Your server probably hasn't been touched since it was created.

```bash
sudo apt update && sudo apt upgrade -y
sudo apt autoremove -y
```

**What this actually does:**

- `apt update` - grabs the latest package info
- `apt upgrade -y` - updates everything (the `-y` just says "yes" to everything)
- `apt autoremove -y` - cleans up leftover packages

## Don't Use Root for Everything

Using root all the time is like driving with your seatbelt off - you'll probably be fine until you're not.

```bash
# Make a new user
sudo adduser adminleo

# Give them sudo powers
sudo usermod -aG sudo adminleo
```

I used Bitwarden to generate a random crazy secure password.

If you want to generate passwords in the terminal (because why not):

```bash
LC_CTYPE=C tr -dc A-Za-z0-9 < /dev/urandom | head -c 32 | xargs
```

## SSH Keys Are Your Friend

Passwords are weak. SSH keys are strong. Here's the deal:

**On your computer:**

```bash
# Make an SSH key
ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519_vps_root

# Lock down the permissions
chmod 600 ~/.ssh/id_ed25519_vps_root
chmod 700 ~/.ssh

# Add it to your SSH agent
eval $(ssh-agent)
ssh-add ~/.ssh/id_ed25519_vps_root
```

**Make connecting easier** by editing `~/.ssh/config`:

```toml
Host hostinger
    HostName 72.xxx.xxx.xxx
    User root
    IdentityFile ~/.ssh/id_ed25519_vps_root
    IdentitiesOnly yes
```

Now you can just type `ssh hostinger` instead of the full command.

**On your server:**

```bash
# Set up the SSH folder
mkdir ~/.ssh

# Add your public key (get this from your local machine)
echo "your-public-key-goes-here" >> ~/.ssh/authorized_keys

# Set permissions (important!)
chmod 700 ~/.ssh
chmod 600 ~/.ssh/authorized_keys
```

**Quick permission explanation:**

- `600` = only you can read/write this file
- `700` = only you can access this folder

## Lock Down SSH

Now we make SSH actually secure:

```bash
sudo vim /etc/ssh/sshd_config
```

**Things to change:**

- `PermitRootLogin prohibit-password` - root can only use SSH keys
- `PasswordAuthentication no` - no more password logins
- `Port 2222` - move away from the default port

**Heads up:** Some cloud providers have their own config files that override this. Check if you have any:

```bash
ls /etc/ssh/sshd_config.d/
sudo vim /etc/ssh/sshd_config.d/50-cloud-init.conf
```

Make sure `PasswordAuthentication no` is set there too.

**Apply the changes:**

```bash
sudo systemctl restart ssh
```

## Set Up a Firewall

UFW (Uncomplicated Firewall) is pretty... uncomplicated:

```bash
# Block everything coming in, allow everything going out
sudo ufw default deny incoming
sudo ufw default allow outgoing

# Allow your SSH port
sudo ufw allow 2222/tcp

# Turn it on
sudo ufw enable
```

**Important:** Test your SSH connection in another terminal before enabling the firewall. Don't lock yourself out.

Reboot after this to make sure everything works:

```bash
sudo reboot
```

## Install Some Useful Stuff

These tools make life easier:

```bash
sudo apt install -y curl wget git vim htop tree unzip fail2ban
```

- `curl/wget` - download things
- `git` - version control
- `vim` - text editor (or use nano if you prefer)
- `htop` - better version of top
- `tree` - see folder structures nicely
- `unzip` - handle zip files
- `fail2ban` - blocks bad actors automatically

## Fail2Ban Setup

This thing monitors your logs and bans IPs that try to break in:

```bash
sudo systemctl enable fail2ban
sudo systemctl start fail2ban

# Check if it's working
sudo systemctl status fail2ban
```

It'll automatically start blocking IPs that try to brute force your SSH.

## Set Your Timezone

Logs with wrong timestamps are annoying:

```bash
sudo timedatectl set-timezone Asia/Manila
```

## Check System Health

Good commands to know:

```bash
# See what's running
htop

# Check disk space
df -h

# Check memory
free -h

# System info with cool ASCII art
neofetch
```

## Basic Monitoring Commands

```bash
# See all listening ports
ss -tuln

# Check system info
uname -a

# See who's logged in
w
```

## What I Actually Learned

- Security isn't optional - every step here prevents real attacks
- SSH keys are way better than passwords
- Firewalls are essential, not optional
- Always test before you lock yourself out
- Document everything (hence this guide)

## What's Next

Now that I have a secure box, I'm thinking about:

- Kubernetes The Hard Way
- Maybe a web server
- Some kind of monitoring

## Quick Commands for Later

```bash
# Updates
sudo apt update && sudo apt upgrade -y

# Add user to sudo group
sudo usermod -aG sudo username

# Restart SSH
sudo systemctl restart ssh

# Firewall status
sudo ufw status

# Check what's listening
ss -tuln
```

That's it. Nothing revolutionary, just a solid foundation that won't get pwned immediately. The internet is full of bots scanning for weak servers, be careful out there.
