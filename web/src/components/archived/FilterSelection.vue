<template>
 <div class="is-inline is-3"
            v-bind:class="{ 'has-background-white': boxOneActive, 'has-background-dark': !boxOneActive}">
        <div class="field">
              <b-switch 
              :value="true" 
              type="is-success"
              v-model="boxOneActive"></b-switch>
          </div>
        <b-field>
              <b-select :disabled="!boxOneActive" placeholder="Select Column" expanded>
                  <option value="flint">Outcome</option>
                  <option value="silver">Age Range</option>
              </b-select>
          </b-field>
          <b-field>
              <b-select :disabled="!boxOneActive" placeholder="Select Value" expanded>
                  <option value="flint">Flint</option>
                  <option value="silver">Arrest</option>
                  <option value="silver">18-24</option>
              </b-select>
          </b-field>
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
      boxOneActive: false
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
