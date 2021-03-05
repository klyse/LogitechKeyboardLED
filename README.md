# Logitech G815 ⌨️ Rider and GoLand Shortcut helper
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fklyse%2FLogitechKeyboardLED.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fklyse%2FLogitechKeyboardLED?ref=badge_shield)


This is a little shortcut helper written in GO 👨‍💻. The software can be started using one of the ``G`` keys. Once started the process profiles every keystroke made on the keyboard. If a certain combination ex. ``SHIFT+CTRL`` is pressed the most important shortcuts for Rider / GoLang are visualized on the keyboard using the RGB key background lightning 💡.

Here's a quick example (sorry for the bad quality, the keys are shining bright and clear in reality🙈):

![](/Images/example.gif)

Here's another one:

![](/Images/example2.gif)

## How to install 💻

1. Clone the repo
2. Build the source code ``go build -o LogitechKeyboardLED.exe ./``  (win only)
3. Navigate to the ``GHub``
4. Add a new launch command with the following params:

![](/Images/GHubConfig.png)

The script checks the state of the process and starts it if the process is stopped, and stops it if the process is running. This way the key works like a toggle.
5. Assign the command to a ``G`` key.
6. Change the working directory in the .ps1 script.
That's it✔️

## How to add shortcuts

The file ``main.go`` has a list of predefined shortcuts. Feel free to add your shortcuts in your repo/fork. I'm always happy if you add more common shortcuts and send them to me via pull-request. :)

# More info

I wrote an article about this repo on [dev.to](https://dev.to/klyse/my-g815-keyboard-is-on-steroids-31fl) have a look :).


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fklyse%2FLogitechKeyboardLED.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fklyse%2FLogitechKeyboardLED?ref=badge_large)