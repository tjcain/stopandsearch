import http from './HTTP';

export default {
    column(category, parameters) {
      return http.get(`/stats/columns/${category}?${parameters}`).then(response => response.data);
    },
    total(parameters) {
      return http.get(`/stats/count?${parameters}`).then(response => response.data)
    }
};