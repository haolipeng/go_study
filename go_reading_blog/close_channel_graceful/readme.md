channel的关闭原则
don’t close a channel from the receiver side and 
don’t close a channel if the channel has multiple concurrent senders.