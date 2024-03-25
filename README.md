Breez SDK demo for Tuscany Lightning Summit 2024

```
% go run . -d data/test
__________                                          ___________                                       
\______   \_______   ____   ____ ________           \__    ___/_ __  ______ ____ _____    ____ ___.__.
 |    |  _/\_  __ \_/ __ \_/ __ \\___   /   ______    |    | |  |  \/  ___// ___\\__  \  /    <   |  |
 |    |   \ |  | \/\  ___/\  ___/ /    /   /_____/    |    | |  |  /\___ \\  \___ / __ \|   |  \___  |
 |______  / |__|    \___  >\___  >_____ \             |____| |____//____  >\___  >____  /___|  / ____|
        \/              \/     \/      \/                               \/     \/     \/     \/\/     
         .____     _______        _________                     .__  __      _______________   ________    _____  
         |    |    \      \      /   _____/__ __  _____   _____ |__|/  |_    \_____  \   _  \  \_____  \  /  |  | 
         |    |    /   |   \     \_____  \|  |  \/     \ /     \|  \   __\    /  ____/  /_\  \  /  ____/ /   |  |_
         |    |___/    |    \    /        \  |  /  Y Y  \  Y Y  \  ||  |     /       \  \_/   \/       \/    ^   /
         |_______ \____|__  /   /_______  /____/|__|_|  /__|_|  /__||__|     \_______ \_____  /\_______ \____   | 
                 \/       \/            \/            \/      \/                     \/     \/         \/    |__| 

sdk> 
```

## What's the plan

Add the Breez SDK to this barebones CLI to connect, receive and send payments. The `final` branch has the completed demo code.

*Add the SDK* - Add some code to access the SDK and listen to SDK events 

*Connect* - Register a node using a Greenlight Invite Code

*Query node state* - Get the node pubkey, on-chain / lightning liquidity

*Receive a payment* - Create an invoice, then receive a payment while opening a channel

*Send a payment* - Pay an invoice using our node
