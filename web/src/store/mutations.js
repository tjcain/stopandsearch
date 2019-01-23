import Vue from 'vue'

export const setCategories = (state, categories) => {
    state.categories = categories
}

export const setURLQueryParameters = (state, {key, parameters}) => {
   Vue.set(state.urlQuery, key, parameters) 
}

export const setSearchData = (state, {key, data}) => {
    Vue.set(state.searchData, key, data)
}

export const setCount = (state, data) => {
    state.totalCount = data
}

export const setLoaded = (state, data) => {
    state.loaded = data
}

export const setFilters = (state, {key, parameters}) => {
    Vue.set(state.filters, key, parameters)
}

export const initFilters = (state, data) => {
    data.map(e => Vue.set(state.filters, e, []))
}