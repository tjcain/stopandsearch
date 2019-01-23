import Environments from '../services/Environment.js';
import Counts from '../services/Counts.js';
import axios from 'axios'

export const fetchCategories = async function({commit})  {
   await Environments.fetchAllCategories()
        .then(data => {commit('setCategories', data)})
    return
}

export const updateURLQueryParameters = ({commit}, {key, parameters}) => {
    commit('setURLQueryParameters', {key, parameters})
}

export const updateFilters = ({commit}, {key, parameters}) => {
    commit('setFilters', {key, parameters})
}

export const fetchCount = ({commit}, parameters) => {
    Counts.total(parameters)
    .then(data => commit('setCount', data))
}

// methods below need attention
export const fetchSearchData = async function({commit, state, getters}) {
    const queryString = getters.getQueryString
    const columns = Object.keys(state.categories)
    
    const promises = columns.map(name => axios.get(`/api/stats/columns/${name}?${queryString}`));
    const responses = await Promise.all(promises)
    columns.forEach(function(key, i){
        let data = responses[i].data
        commit('setSearchData', {key, data})
    })
    commit('setLoaded', true)
}

// export const fetchSearchData = async function({commit, state, getters}) {
//         const queryString = getters.getQueryString
//         const columns = Object.keys(state.categories)
        
//         const promises = columns.map(name => fetch(`/api/stats/columns/${name}?${queryString}`)
//             .then(r => r.json()));
//         const responses = await Promise.all(promises)
//         columns.forEach(function(key, i){
//             let data = responses[i]
//             commit('setSearchData', {key, data})
//         })
//         commit('setLoaded', true)
// }


