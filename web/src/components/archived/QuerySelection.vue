<template>
    <div class="columns is-block ">
        <!-- <div class="column">
            <h5 class="title is-5">2) Create a query</h5>
        </div> -->

        <div class="column is-vertical-center">
            <p> I am interested in examining &nbsp; </p>

            <b-dropdown v-model="firstColumn">
                <button class="button is-primary" slot="trigger" type="button">
                    <span>{{firstColumn}}</span>
                    <b-icon icon="menu-down"></b-icon>
                </button>
                
                <b-dropdown-item 
                    v-for="option in options" 
                    :key=option.id
                    :value="pluralize(option.name)"
                    class="is-success">
                        {{option.name}}
                </b-dropdown-item>
            </b-dropdown>

        <div class="is-inline">
            
        </div>

        </div>
        <!-- <span> &nbsp; where &nbsp; is something eh</span> -->
       
        </div>
</template>

<script>
import sourceData from "@/data.json";
import DropDownButton from "@/components/DropDownButton.vue";

import pluralize from "pluralize";
import vueSlider from "vue-slider-component";

import slideBarSettings from "@/slideBarSettings.js";

export default {
  data() {
    return {
      options: sourceData.data,
      forces: sourceData.forces,
      firstColumn: "Column",
      secondColumn: "Column",
      secondOption: "Option",
      slideBarSettings: slideBarSettings,
      forceSelection: "All Police Forces"
      //   columns: ["Outcomes", "Ethnicities", "Age groups", "Objects of search"]
    };
  },
  methods: {
    pluralize
  },
  components: { DropDownButton, vueSlider }
};
</script>

<style scoped>
.is-vertical-center {
  display: flex;
  align-items: center;
}

.has-padding-top {
  padding-top: 1vh;
}

.pd-bt {
  padding-bottom: 30px;
}

.bg {
  background: red;
}
</style>
