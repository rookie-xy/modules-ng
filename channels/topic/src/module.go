package topic

import . "github.com/rookie-xy/worker/types"

var topicModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    topicCommands,
    CHANNEL_MODULE,
}

var channelTopicModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    channelTopicContext,
    channelTopicCommands,
    TOPIC_MODULE,
}

func init() {
    Modules = append(Modules, &Topic{Module_t:topicModule},
                              &ChannelTopic{Module_t:channelTopicModule})
}
