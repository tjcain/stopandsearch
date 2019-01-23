export const getQueryString = (state) => Object.values(state.urlQuery).join("&")

export const getAgeRangeData = (state) => state.searchData["age_range"]
export const getEthnicityData= (state) => state.searchData["ethnicity"]
export const getOutcomeData = (state) => state.searchData["outcome"]
export const getObjectOfSearchData = (state) => state.searchData["object_of_search"]
export const formatCount = (state) => String(state.totalCount).replace(/(.)(?=(\d{3})+$)/g,'$1,')

// These need re-doing... generate in back end...
export const getPositiveOutcomePercent = (state) => {
    const outcomes = state.searchData["outcome"]
    let n = 0
    outcomes.forEach(e => {
        if (e.name !== "Nothing Found") {
            n += e.count
        }
    });
    return Math.round((n / state.totalCount) * 100)
}

// bad
export const getNoOutcomePercent = (state) => {
    const outcomes = state.searchData["outcome"]
    let n = 0
    outcomes.forEach(e => {
        if (e.name === "Nothing Found") {
            n += e.count
        }
    });
    return Math.round((n / state.totalCount) * 100)
}

// bad
export const getArrestPercentage = (state) => {
    const outcomes = state.searchData["outcome"]
    let n = 0
    outcomes.forEach(e => {
        if (e.name === "Arrest") {
            n += e.count
        }
    });
    return Math.round((n / state.totalCount) * 100)
}