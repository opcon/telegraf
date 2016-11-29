#!/bin/bash

if ! grep "^telegraf:" /etc/group &>/dev/null; then
    groupadd -r telegraf
fi

if ! id telegraf &>/dev/null; then
        if useradd -h 2>&1 | grep -q "^[[:space:]]*-M"; then
        # Newer version of useradd have -M option to force not creating home dir
        useradd -r -M telegraf -s /bin/false -d /etc/telegraf
    else
        # Older version does not create home dir by default
        useradd -r telegraf -s /bin/false -d /etc/telegraf
    fi
fi

if [[ -d /etc/opt/telegraf ]]; then
    # Legacy configuration found
    if [[ ! -d /etc/telegraf ]]; then
        # New configuration does not exist, move legacy configuration to new location
        echo -e "Please note, Telegraf's configuration is now located at '/etc/telegraf' (previously '/etc/opt/telegraf')."
        mv -vn /etc/opt/telegraf /etc/telegraf

        if [[ -f /etc/telegraf/telegraf.conf ]]; then
            backup_name="telegraf.conf.$(date +%s).backup"
            echo "A backup of your current configuration can be found at: /etc/telegraf/${backup_name}"
            cp -a "/etc/telegraf/telegraf.conf" "/etc/telegraf/${backup_name}"
        fi
    fi
fi
