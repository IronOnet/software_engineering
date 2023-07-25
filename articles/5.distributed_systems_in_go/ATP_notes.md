## The Authenticated Transfer Protocol 

The authenticated transfer protocol aka ATP is a protocol for large scale 
distributed applications

### Identity 

Users are identified by domain names in AT protocol. These domains map to  
cryptographic URLs which secure the user's account and its data 

===> @alice.com             Domain names 
===> at://alice.com         URLS 
===> at://did:plc:123..yz/  Cryptographic URLS

### Data Repositories 

User data is exchanged in signed data repositories. these repositories 
are collections of records which include posts, comments, likes, follows, 
media blobs. etc 


E.g : @alice.come =============> signs ==============> repo records

### Federation 

ATP syncs the repositories in a federated networking model. Federation was chosen 
to ensure the network is convenient to use and reliably available. 
Commands are sent between servers using HTTPS + XRPC

### InterOperation 

A global schemas network called lexicon is used to unify the names and behaviors 
of calls accross the servers. Servers implement "lexicons" to support 
featureset, including the core ATP Lexicon for syncing user repositories 
and the Bsky Lexicon to provide basic social behaviors 

The AT Protocol exchanges schematic and semantic information, enabling the 
software from different orgs to understand each others data. This gives 
ATP clients freedom to produce user interfaces independently of the servers 
and remove the need to exchange rendering code (HTML/JS/CSS) while browsing content 


### Achieving Scale 

ATP distinguishes between "small-world" vs "big-world" networking. Small-world 
networking encompasses inter-personal activity whil big-world networking 
aggregates activity outside of the user's personal interactions. 

    - Small-world: delivery of events targeted at specific users such as mentions, 
    replies, and DMs, and sync of datasets according to follow graphs. 

    - Big-world: large-scale metrics (like, reposts, followerws), content discovery 
    (algorithms), and user search. 

==> Personal data servers are responsible for small-world networking while  indexing services separately crawl the network to provide Big-world networking.

==> The small-world/big-world distinction is intended to achieve scale as well as a high degree of 
user choice.


### Algorithmic choice 

As with web search engines, users are free to select their indexers. Each feed discovery 
section, or search interface is integrated into the PDS while being served from a third party 
service. 


                    ===============> @news.com  Breaking Stories 

            ==> PDS          ==========> @tech-community.com Top Posts 

                    ===============> @cats.com Best cats daily


### Account portability 

A personal data server might fail at any time, either by going offline entirely
or ceasing service for specific users. ATP goal is to ensure that a user can 
migrate their account to a new PDS without server involvement. 

User data is stored in signed data repositories and verified by DIDs (Decentralized Identifiers). 
DIDs are essentialy registries of user certificates, similar in some ways 
to the TLS certificate system. They are expected to be secure, reliable and 
indpendent of user's PDS

Each DID document publishes two public keys: a singing key and a recovery key. 

    - SInging key: Asserts changes to the DID document and the user's data repository 
    - Recovery key: Asserts changes to the DID document, may override the singing key 
    within a 72 hours window. 

The singing key is entrusted to the PDS so that it can manage user's data, but the recovery 
key is saved by the user, e.g as a paper key this makes it possible for the user to 
update their account to a new PDS without the original host help. 

A backup of the user's data is persistently synced to their client as a backup 
(contigent on the disk space available). Should a PDS disapear without notice, 
the user should be able to migrate to a new provider by updating their DID Document 
and uploading the backup 