package elasticsearch

import . "github.com/rookie-xy/worker/types"

var elasticSearchModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    elasticSearchCommands,
    OUTPUT_MODULE,
}

var outputElasticSearchModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    outputElasticSearchContext,
    outputElasticSearchCommands,
    ELASTICSEARCH_MODULE,
}

func init() {
    Modules = append(Modules, &ElasticSearch{Module_t:elasticSearchModule},
                              &ElasticSearchOutput{Module_t:outputElasticSearchModule})
}
