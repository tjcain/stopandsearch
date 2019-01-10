/* 
 * Mocking client-server processing
 */

// search counts is a mock of the data that will be returned by the api in one call
// with the provided parameters. The categories have come from the stop and search
// reports from the data.police.uk website. 
const _searchCounts = {
    count: 1234,
    positiveOutcome: 234,
    noOutcome: 1000,
    ageRange: [{
        "Under 10": 5
    }, {
        "10-17": 20
    }, {
        "18-24": 50
    }, {
        "25-34": 20
    }, {
        "Over 34": 5
    }],

    outcome: [{
        "No further action": 100
    }, {
        "Arrest": 20
    }, {
        "Khat / Cannabis": 10
    }, {
        "Warning": 12
    }, {
        "Community Resolution": 18
    }, {
        "Summons Penalty": 29
    }, {
        "Notice for Disorder": 44
    }, {
        "Caution": 32
    }],

    objectOfSearch: [{
            "Drugs": 29
        },
        {
            "Stolen property": 21
        },
        {
            "Firearms": 17
        },
        {
            "Offensive weapons": 14
        },
        {
            "Criminal damage": 7
        },
        {
            "Going equipped": 10
        },
        {
            "Other": 8
        },
    ],

    selfDefinedEthnicity: [{
            "White": 100
        },
        {
            "Black": 50
        },
        {
            "Asian": 20
        },
        {
            "Chinese or other": 40
        },
        {
            "Mixed": 10
        },
        {
            "Not Stated": 40
        },
    ],



}

export default {
    getSearches(cb) {
        setTimeout(() => cb(_searchCounts), 100)
    }
}