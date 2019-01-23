<template>
  <div class="box noselect">
    <div>
      <nav class="level is-mobile">
        <div class="level-left" @click="toggleFilterList">
          <div class="level-item">
              <span class="icon is-small has-text-primary" 
                    v-bind:class="[showFilterList ? 'rotated' : '']">
                <i class="fas fa-chevron-down"></i>
              </span>
        
          </div>
          <div class="level-item">
            <p class="menu-label">{{title.split("_").join(" ")}}</p>
          </div>
          <div class="level-item filler">
            <!-- filler -->
          </div>
        </div>

        <div class="level-right" @click="active = !active">
          <div class="level-item">
            <b-switch size="is-small" v-model="active" :value="true" type="is-success"></b-switch>
          </div>
        </div>
      </nav>
      <app-drop-down-filter-list
        v-if="showFilterList"
        :active="active"
        :category="title"
        :filters="filters"
      />
    </div>
  </div>
</template>

<script>
import AppDropDownFilterList from "@/components/AppDropDownFilterList";
export default {
  components: {
    AppDropDownFilterList
  },
  data() {
    return {
      showFilterList: false,
      active: false
    };
  },
  computed: {
    getQueryString() {
      return this.$store.getters.getQueryString;
    },
    filterList() {
      return this.$store.state.filters[this.title];
    }
  },
  props: ["title", "filters"],
  methods: {
    toggleFilterList() {
      this.showFilterList = !this.showFilterList;
    }
  },
  watch: {
    active() {
      if (this.active === false) {
        const payload = { key: this.title, parameters: [] };
        this.$store.dispatch("updateURLQueryParameters", payload);
      } else {
        const payload = {
          key: this.title,
          parameters: this.filterList.join("&")
        };
        this.$store.dispatch("updateURLQueryParameters", payload);
      }
    }
  }
};
</script>

<style css scoped>
@media only screen and (min-device-width: 1024px) {
  .menu-label {
    max-width: 10vw;
  }
}

.noselect {
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
}

.level-left, .level-right {
  cursor: pointer;
}

.box {
  padding: 10px;
}

ul {
  padding: 0 10px;
}

.rotated {
  animation: rotate 3s infinite;
  transform: rotate(180deg);
}
</style>
