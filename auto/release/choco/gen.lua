sh("choco pack")
-- choco apikey --key=YOUR_API_KEY --source=https://push.chocolatey.org/
sh("choco push ./yock.nupkg --source=https://push.chocolatey.org/")
