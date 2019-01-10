<template>
    <div class="columns is-block">
        <!-- <div class="column">
            <h5 class="title is-5">1) Select Parameters</h5>
        </div> -->

        <div class="column">
            <div class="columns has-padding-top">
                <div class="column is-7">
                   <p class="pd-bt"><strong>Date Range:</strong></p> 
                   <vue-slider ref="slider4" v-bind="slideBarSettings" v-model="slideBarSettings.value"></vue-slider>
                </div>
                <div class="column">
                    <p class="pd-bt"><strong>Police Force:</strong></p>
                    <b-dropdown v-model="forceSelection">
                        <button class="button is-primary" slot="trigger">
                            <span>{{forceSelection}}</span>
                            <b-icon icon="menu-down"></b-icon>
                        </button>
                        <b-dropdown-item 
                            v-for="force in forces" 
                            :key=force.id
                            :value="force.name"
                            class="is-success">
                                {{force.name}}
                        </b-dropdown-item>
                    </b-dropdown>
                </div>
            </div>
            </div>
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
