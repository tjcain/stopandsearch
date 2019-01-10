<template>
    <b-dropdown v-model="firstColumn">
            <button class="button is-primary" slot="trigger" type="button">
                <span>{{span}}</span>
                <b-icon icon="menu-down"></b-icon>
            </button>
            
            <b-dropdown-item 
                v-for="option in this.options" 
                :key=option.id
                :value="pluralize(option.name)"
                class="is-success">
                    {{option.name}}
            </b-dropdown-item>
    </b-dropdown>
</template>

<script>
import pluralize from "pluralize";
export default {
  props: ["options", "startSpan"],
  methods: {
    pluralize
  },
  data() {
    return {
      span: this.startSpan
    };
  }
};
</script>

<style>
</style>
