# rvspub

In the early 2000s, a cheat client was deliberately leaked for QuakeWorld. The
client phoned home, and data was collected for a year before the list of
cheaters was made public.

This is an attempt to reverse engineer the data format used to store the player
information.

The database contains `230 232` unique entries. The first entry was inserted
on `2000-01-26 at 23:44:02 +0000 UTC`, and the most recent entry is from
`2001-01-29 at 08:49:54 +0000 UTC.`

If you have any information, perhaps the cheat client or any other details
that might be useful, please reach out to me!

## Caveat

The presence of a name in this database does not definitively prove that the
individual behind that name was cheating. It only indicates that the leaked
cheat client was run by someone using that nickname at some point.

## Symbols

I don't have access to the cheat client and haven't been able to determine
with 100% certainty what the different symbols mean, but here's my best guess:

- `A!` - aimbot
- `B!` - wallhack
- `E!` - modified eyes model
- `F!` – fullbright
- `P?` - modified player model
- `R!` - radar
- `S!` - speed cheat

## Usage

```sh
$ make
$ ./rvspub -data-file reference/rvs.dat | head
version,external_ip_address,internal_ip_address,session_id,hostname,signature,timestamp,duration,server_address,server_hostname,user_id,name,eyes_model,player_model,frags,cl_accel,cheats,is_spectator
2.1,24.30.107.69,24.30.107.69,8e8c68e4,screamer.we.mediaone.net,45e4b0253a4ff5082a0421b39291a5f6,2000-01-26T23:44:02Z,2000-01-26T23:44:42Z,209.52.6.252:27501,SmarttNet's Custom MegaTF Server,0,-=Gaspar=-,6967,13845,0,100,R!,false
2.1,207.172.164.7,207.172.164.7,8f91b6f8,pavilion,45e4b0253a4ff5082a0421b39291a5f6,2000-01-27T00:09:04Z,2000-01-27T00:09:04Z,209.246.197.66:27500,[GOD].GameTrak.#2.CA,0,lasdf,6967,33168,0,100,P?,false
2.1,209.103.43.118,209.103.43.118,9e94a87b,anthony,45e4b0253a4ff5082a0421b39291a5f6,2000-01-27T00:00:34Z,2000-01-27T00:11:41Z,24.112.63.188:27500,Clan BEL MEGATF Public Server,0,Zeus,6967,13845,3,400,S! R!,false
2.1,207.172.164.7,207.172.164.7,8f91b6f8,pavilion,45e4b0253a4ff5082a0421b39291a5f6,2000-01-27T00:09:36Z,2000-01-27T00:12:43Z,209.246.197.66:27500,[GOD].GameTrak.#2.CA,0,lasdf,6967,33168,0,100,P? R!,false
2.1,216.209.204.4,216.209.204.4,9594330f,jon,45e4b0253a4ff5082a0421b39291a5f6,2000-01-27T00:19:05Z,2000-01-27T00:20:00Z,24.112.105.208:27500,BJServerOfLove,0,{e}Jon,6967,33168,0,100,P? R!,false
2.1,63.23.94.65,63.23.94.65,8b95d408,thomas-hwee,45e4b0253a4ff5082a0421b39291a5f6,2000-01-27T00:19:24Z,2000-01-27T00:20:19Z,208.195.192.5:27500,[I'm 12]'s Jungle Gym: QF or vanilla,0,newbie,6967,33168,0,9999,P? S! R!,false
2.1,203.109.137.46,203.109.137.46,969358ae,ihug.ihug.co.nz,45e4b0253a4ff5082a0421b39291a5f6,2000-01-27T00:37:45Z,2000-01-27T00:39:19Z,games.gen.nz,You Abuse and you lose,0,,6967,33168,0,100,P? R!,false
2.1,158.252.248.217,158.252.248.217,938a71bc,lyle,45e4b0253a4ff5082a0421b39291a5f6,2000-01-27T00:36:34Z,2000-01-27T00:41:15Z,localhost,Intel TF2.9 Clan[GT],0,Glint,0,0,-2,100,E! P?,false
2.1,63.26.145.242,63.26.145.242,9b964532,default,45e4b0253a4ff5082a0421b39291a5f6,2000-01-27T00:59:53Z,2000-01-27T01:17:03Z,static.com:27500,Static's Clan Arena Damage,0,practiceing name,10539,13845,5,120,E! S!,false
```

## Screenshots

The following are screenshots from the original `rvs13.exe`, which was published
together with the `dat` file back in 2001.

![Player Search](./resources/playersearch.jpg)

![Player Information](./resources/playerinformation.jpg)

## Disclaimer

The original `.exe` file that was released alongside the database can be found
in the `resources` directory. I have not yet analyzed the binary, so run it at
your own risk!

The dataset that was published is also available in the `resources` directory.
This is the database I have been using as a reference while writing this tool.

## Credits

Special thanks to the following individuals for providing valuable data and
information:

- HangTime
- ParadokS
- splash
- To the author(s) of RVS, whoever you may be
- Novate
- RoT// Red Fox

## Final Note

CRIME DOSE NOT PAY [sic]
