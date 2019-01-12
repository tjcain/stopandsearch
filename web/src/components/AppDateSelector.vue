<template>
  <div>
    <p class="menu-label">Date Range</p>
    <div class="slider">
      <div class="field">
        <vue-slider
          ref="slider4"
          v-bind="slideBarSettings"
          v-model="sliderModel"
          v-on:drag-end="$emit('date-range', generateDateParams())"
        ></vue-slider>
      </div>
    </div>
  </div>
</template>

<script>
import slideBarSettings from "@/slideBarSettings.js";
import vueSlider from "vue-slider-component";
export default {
    components: {
    vueSlider
  },
  methods: {
    generateDateParams: function() {
      return this.dateValue
    },
  },
  computed: {
    dateValue() {
      let from = this.sliderModel[0].split("-")
      from = from[0] + "-01-" + from[1]
      let to = this.sliderModel[1].split("-")
      to = to[0] + "-01-" + to[1]
      return "time="+from+"&"+"time="+to
    }
  },
  data() {
    return {
      slideBarSettings: slideBarSettings,
      sliderModel: ["09-2015", "01-2019"],
    };
  },
  created: function () {
      this.$emit('date-range', this.dateValue)
  }
};
</script>

<style scoped>
.field:not(:last-child) {
  margin-bottom: 1px;
}
.field {
  margin-bottom: 1px;
}
.checkboxes {
  /* height: 50px; */
  overflow-y: scroll;
}
.slider {
  width: 100%;
  padding-top: 20px;
}
</style>

