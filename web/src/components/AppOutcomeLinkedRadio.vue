<template>
    <section>
        <p class="menu-label">Outcome Linked to Object of Search</p>
        <div class="field">
            <b-radio 
                @input="send"
                v-model="radio"
                native-value="true"
                size="is-small">
                Yes
            </b-radio>
        </div>
        <div class="field">
            <b-radio v-model="radio"
                @input="send"
                native-value="false"
                size="is-small">
                No
            </b-radio>
        </div>
        <div class="field">
            <b-radio v-model="radio"
                @input="send"
                native-value="neither"
                size="is-small">
                Both
            </b-radio>
        </div>                
    </section>
</template>

<script>
export default {
    data() {
        return {
            radio: "neither"
        }
    },
    computed: {
        queryParam() {
            if (this.radio === "neither") {
                return ""
            }
            return "outcome_linked_to_object=" + this.radio
        }
    },
    methods: {
        send() {
            console.log("test")
            this.$emit('update-radio', this.queryParam)
        }
    },
    created: function () {
        this.$emit('update-radio', this.queryParam)
    }
}
</script>

<style scoped>
    .field:last-child {
    margin-bottom: 1px;
    }
    .field {
    margin: 1px 0px;
    }
    .checkboxes {
    max-height: 125px;
    overflow-y: scroll;
    }
    .menu-label {
    margin: 1px 0
    }
</style>
