import http from './HTTP';

export default {
    // fetchAllCategories will return all of the available category information
    // which allows us to render the filter options
    fetchAllCategories() {
      return http.get(`/categories`).then(response => response.data);
    },
    // Add other things such as version
  };