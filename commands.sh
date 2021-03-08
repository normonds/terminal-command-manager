# ls
ls -lpah --color | awk '{k=0;for(i=0;i<=8;i++)k+=((substr($1,i+2,1)~/[rwx]/)*2^(8-i));if(k)printf("%0o ",k);print}'
cat /proc/cpuinfo
# info linux
printf 'whoami: ';whoami;printf 'uname -a: ';uname -a;lsb_release -a;lscpu | grep "CPU(s):\\|Model name\\|per socket\\|CPU MHz\\|Vendor ID";awk '$3=="kB"{if ($2>1024^2){$2=$2/1024^2;$3="GB";} else if ($2>1024){$2=$2/1024;$3="MB";}} 1' /proc/meminfo | column -t | grep "MemTotal\\|MemFree\\|SwapTotal\\|SwapFree"

# ps linux
ps -afx
# top linux
echo "press V tx2 mx2";top -c -E g
# linux release info
lsb_release -a

gcloud components update
gcloud config list && echo;echo;gcloud app describe
gcloud projects list

ibmcloud update; ibmcloud plugin update
# make active nvm node sudoable
n=$(which node); n=${n%/bin/node}; chmod -R 755 $n/bin/*; sudo cp -r $n/{bin,lib,share} /usr/local

# kubernetes ibmcloud echo kube config
kubernetes_config=$(ibmcloud ks cluster-config nooni-cluster --export) && echo;echo $kubernetes_config;echo;
# kubernetes ibmcloud info
echo --- NODES ---;kubectl get nodes;echo --- PODS ---;kubectl get pods;echo --- SERVICES ---; kubectl get services
# kubernetes ibmcloud workers info
ibmcloud ks workers --cluster <prompt:cluster id:bku5rfnf0tefg7alcca0>

# kubernetes conn terminal
pods=$(kubectl get pods -o name) && pod=$(eval echo $pods | cut -d'/' -f 2) && echo kubectl exec -it $pod -- top && eval kubectl exec -it $pod -- byobu
# l linux filesystem list al files in directory
ls -ahsXp --color
# ls linux filesystem list files permissions
ls -lpah --color | awk '{k=0;for(i=0;i<=8;i++)k+=((substr($1,i+2,1)~/[rwx]/)*2^(8-i));if(k)printf("%0o ",k);print}'
# linux format volume sudo
mkfs -t ext4 /dev/someVolume

# linux filesystem size disks sudo
fdisk -l
# linux filesystem size disks
df -h
# linux filesystem size disks
lsblk -o NAME,FSTYPE,LABEL,SIZE,UUID,MOUNTPOINT
# linux filesystem files in current directory
find -type f | wc -l
# linux filesystem current directory size
du -hs <prompt:directory:.>
# linux app show largest files
ncdu -x <prompt:directory:/>
# ports listen linux opened
netstat -plnt | grep LISTEN --color=always
# linux opened ports sudo
lsof -i -P -n | grep LISTEN --color=always
# linux system info
lshw -short
# linux system info cpu
lscpu
# linux current directory
pwd
# linux firewall show iptables sudo
iptables -L
# linux edit sudo (@reboot /path/to/script.sh)
crontab -e
# linux firewall accept all connections for session sudo
iptables -P INPUT ACCEPT && iptables -P OUTPUT ACCEPT && iptables -P FORWARD ACCEPT && iptables -F
# linux systemctl service status
systemctl status <prompt:service name:cron>
# chmod info
echo "7:rwx,6:rw-,5:r-x,4:r--,3:-wx,2:-w-,1:--x,0---"

#imagemagick resize pngs to out folder
mogrify -monitor -path out -resize 33% -format png *.png
#imagemagick convert jpgs to pngs imgs
mogrify -format png *.jpg
# ip info url request
curl https://ifconfig.co
# ipjson info url request
curl https://ifconfig.co/json
# port info url request
curl https://ifconfig.co/port/<prompt:port:8080>
# nodejs lts install sudo
curl -sL https://deb.nodesource.com/setup_<prompt:LTS version 10,12,14:14>.x | sudo -E bash - && sudo apt-get install nodejs
# grep linux search string in files nonbinary recursive caseinsensitive
grep -rnIi --include \\<prompt:filetype:*.*> "<prompt:search string>" <prompt:directory:.> --color=always | more
# find linux search file
find <prompt:directory:/home> -iname "*<prompt:string in filenamepath:json>*" -print 2>/dev/null | grep <prompt:string in filenamepath:json> -i --color=always | more
# linux append
echo "<prompt:string>" >> <prompt:filename:file.txt>
# linux replace text file contents
sed -i "s/<prompt:search string>/<prompt:replace string>/g" <prompt:filename:file.txt>
# linux show first lines of file
head -<prompt:first n lines:30> <prompt:file:file.txt>

mysqldump --user=<prompt:user:root> --password=<prompt:password:1234> <prompt:database> --result-file <prompt:file to save>
# linux
zip <prompt:zipped file:file.zip> <prompt:file to zip:file.txt>

# git history
git log --graph --decorate --pretty=oneline --abbrev-commit
# github publish
git push -u origin master

# npm show glob packages
npm list -g --depth 0
# npm show outdated glob packages
npm outdated -g --depth=0

# package info
apt-cache show <prompt:package name:google-chrome-stable>

# youtube 1 fps
ffmpeg -loop 1 -framerate 1 -i <prompt:image file> -i <prompt:mp3 file> -c:v libx264 -preset veryslow -crf 0 -c:a copy -shortest <prompt:output file:output.mp4>

# editor open pipe
apt-cache show nano | nano -
# editor
echo "nano: F1 - help;SHIFT+ALT+4 - toogle word wrap; ALT+U - undo; ALT+N - redo;CTRL+^ - start text mark; CTRL+K - cuts selected text;CTRL+U - paste;F6 - search; ALT+W - repeat search;ALT+C - toogle info box"
