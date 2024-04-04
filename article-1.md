
# Improving visibility of my finances

I have being struggling to keep track of my finances, in the sense of knowing with what I'm spending my money,
what are my debts, etc. This is the basics to achieve a good balance of spending and saving, it's how you start
planning future decisions.

The current state of solutions I know struggles alot. I have tried to manually keep track of my bills using spreadcheats or something similar which worked in the beginning but at the moment I skip a single entry, I compromised my commitment with this task.

My dream tool is something with a great level of automation which tracks my spenses from all (most) of my bank accounts. I want to know the category of spenses (foo, entertainment, bills, etc), when it happend, the amount, and some sum of categories or something. This will give me visibility of how my spenses are distributed.

# My goal

The goal is to find a way to configure transaction notifications on the bank app in a way where I can store this information and expose a report of the current state when needed.

It would be perfect if the banks I use offer this type of notification feature but it didn't. The only way I could bypass this limitation was that in one of the banks it's possible to configure e-mail notifications for transactions. BOOM I just need a way to watch and process e-mails.

## Getting access to e-mails

I started thinking on how can one have access to e-mails sent to an specific address. I haven't played a lot with e-mail protocols on my carreer but I knew SMTP and that was my first tought. I've searched "reading e-mails with smtp" but later discovered I was looking for the wrong protocol, what I need is the `IMAP`.

IMAP stands for `Internet Message Access Protocol` (I'm not going to mention POP protocol here). Basically one's e-mails are stored on a server, if you use gmail.com your e-mails are stored on a server at Google. But when we wan't to have access to that e-mails on some kind of client the process is not so simple as just downloading them. What IMAP does is to download a linked copy of the e-mails, which means anything we do locally using IMAP is applied to the server. That's how we can read a new e-mail on the cellphone and it will appear marked as read on the computer.

# The desing

I will have a `daemon` that connects to the `IMAP` server 

