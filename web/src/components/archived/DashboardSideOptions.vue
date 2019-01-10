<template>
  <aside class="menu is-hidden-mobile">
    <!-- FORCE SELECTOR -->
    <p class="menu-label">Police Force</p>
    <div class="dropdown">
      <div class="field">
        <select-force/>
      </div>
    </div>

    <!-- SLIDER -->
    <p class="menu-label">Date Range</p>
    <div class="slider">
      <div class="field">
        <vue-slider
          ref="slider4"
          v-bind="slideBarSettings"
          v-model="sliderModel"
          v-on:drag-end="$emit('params', generateGroupParams())"
        ></vue-slider>
      </div>
    </div>

    <!-- AGE RANGE -->
    <p class="menu-label">Age Range</p>
    <section class="checkboxes">
      <div v-for="age in checkBoxValues.ageRange" :key="age.v" class="field">
        <b-checkbox
          v-on:input="$emit('params', generateGroupParams())"
          size="is-small"
          v-model="groupParams"
          :native-value="age.v"
        >{{age.l}}</b-checkbox>
      </div>
    </section>

    <!-- OBJECT OF SEARCH -->
    <!-- <p class="menu-label">Object of search</p>
    <section class="checkboxes">
      <div v-for="object in objectOfSearchGroup" :key="object" class="field">
        <b-checkbox size="is-small" v-model="objectOfSearchGroup" :native-value="object">{{object}}</b-checkbox>
      </div>
    </section>-->
  </aside>
</template>

<script>
import SelectForce from "@/components/SelectForce.vue";

// date slider
import vueSlider from "vue-slider-component";
import slideBarSettings from "@/slideBarSettings.js";
export default {
  components: {
    SelectForce,
    vueSlider
  },
  methods: {
    generateGroupParams: function() {
      const time =
        "&time=" + this.sliderModel[0] + "&time=" + this.sliderModel[1];
      let v = this.groupParams.join("&");
      v = v + time;
      console.log(v);
      return v;
    }
  },
  data() {
    return {
      count: 0,
      slideBarSettings: slideBarSettings,
      sliderModel: ["09-2015", "12-2018"],
      checkBoxValues: {
        ageRange: [
          { l: "Under 10", v: "age_range=under+10" },
          { l: "10 to 17", v: "age_range=10-17" },
          { l: "18 to 24", v: "age_range=18-24" },
          { l: "25 to 34", v: "age_range=25-34" },
          { l: "Over 34", v: "age_range=over+34" }
        ],
        outcome: [
          { l: "Suspect arrested" },
          { l: "Penalty Notice for Disorder" },
          { l: "Offender cautioned" },
          { l: "Khat or Cannabis warning" },
          { l: "Community resolution" },
          { l: "Offender given penalty notice" },
          { l: "Arrest" },
          { l: "Caution (simple or conditional)" },
          { l: "Local resolution" },
          { l: "" },
          { l: "Offender given drugs possession warning" },
          { l: "Article found - Detailed outcome unavailable" },
          { l: "Suspect summonsed to court" },
          { l: "Summons / charged by post" },
          { l: "A no further action disposal" }
        ]
      },

      groupParams: [
        // AGE_RANGE
        "age_range=under+10",
        "age_range=10-17",
        "age_range=18-24",
        "age_range=25-34",
        "age_range=over+34"
        // ETHNICITY
      ]
    };
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

/* CSS BELOW MAKES SCROLLBAR SHOW  */
/* .checkboxes::-webkit-scrollbar {
  -webkit-appearance: none;
}
.checkboxes::-webkit-scrollbar:vertical {
  width: 11px;
}
.checkboxes::-webkit-scrollbar:horizontal {
  height: 11px;
}
.checkboxes::-webkit-scrollbar-thumb {
  border-radius: 8px;
  border: 2px solid white; /* should match background, can't be transparent 
  background-color: rgba(0, 0, 0, 0.5); 
} */
/* **************** */
</style>
