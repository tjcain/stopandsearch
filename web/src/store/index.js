import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import * as mutations from './mutations';
import * as actions from './actions';
import * as getters from './getters';

export default new Vuex.Store({
  state: {
    categories: {},
    // todo: auto generate this
    filters: {
      "age_range": [],
      "ethnicity": [],
      "outcome": [],
      "outcome_linked_to_object": [],
      "object_of_search": [],
      "legislation": []
    },
    urlQuery: {},
    searchData: {},
    totalCount: 100,
    loaded: false
    // consider an 'active columns' list that the filter toggles control
    // then loop through that to get the filter settings? 
  },
  getters,
  mutations,
  actions,
})
