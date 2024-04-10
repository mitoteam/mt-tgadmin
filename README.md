# mt-tgadmin

[![Go Report Card](https://goreportcard.com/badge/github.com/mitoteam/mt-tgadmin)](https://goreportcard.com/report/github.com/mitoteam/mt-tgadmin)
[![GitHub](https://img.shields.io/github/license/mitoteam/mt-tgadmin)](https://github.com/mitoteam/mt-tgadmin/blob/main/LICENSE)

[![GitHub Version](https://img.shields.io/github/v/release/mitoteam/mt-tgadmin?logo=github)](https://github.com/mitoteam/mt-tgadmin)
[![GitHub Release Date](https://img.shields.io/github/release-date/mitoteam/mt-tgadmin)](https://github.com/mitoteam/mt-tgadmin/releases)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/mitoteam/mt-tgadmin)
[![GitHub contributors](https://img.shields.io/github/contributors-anon/mitoteam/mt-tgadmin)](https://github.com/mitoteam/mt-tgadmin/graphs/contributors)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/mitoteam/mt-tgadmin)](https://github.com/mitoteam/mt-tgadmin/commits)
[![GitHub downloads](https://img.shields.io/github/downloads/mitoteam/mt-tgadmin/total)](https://github.com/mitoteam/mt-tgadmin/releases)

Simple GUI to send messages from Telegram bot to send to group, channel or supergroup. Possible to send messages anonymously as admin.

## How To

We try to keep `mt-tgadmin help` accurate.

Before setting up `mt-tgadmin` you need to register your bot with **@BotFather**, add bot to desired group and obtain this group's **chatID**.

* Registering new bot with [@BotFather](https://t.me/botfather) (useful [How To](https://core.telegram.org/bots/features#creating-a-new-bot))
* Add created bot to your group.
* Obtaining group's **chatID**: Open [@RawDataBot](https://t.me/rawdatabot) and follow it instructions.

### Install

You will need bot's **token** and **chatID** to run `mt-tgadmin`.

* Unpack archive for you platform to desired location (there is just one executable file in archive).
* Run `mt-tgadmin init` to create simple config to start with.
* Open created file `.bot.yml` with text editor and edit settings. Example settings are in [.bot.EXAMPLE.yml](https://github.com/mitoteam/mt-tgadmin/blob/main/.bot.EXAMPLE.yml).
* Run `mt-tgadmin run` to check the setup. You should be able to open WebGUI at this point. Press `Ctrl + C` to stop it if everything is OK.
* Run `mt-tgadmin install` to register it as daemon.
* Consider using nginx or other webserver as reverse proxy to use TLS (HTTP**S**) and other extended http features.

### Upgrade

* Just unpack newer version binary and restart process (or service).
