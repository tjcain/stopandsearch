// I hate this component..  Doing everything the long way is the only way I
// could get anything to work.. I'm sorry.
<template>
    <div>
        <p class="menu-label">Age Range</p>
        <section class="checkboxes">
            <div class="field">
                <b-checkbox 
                @input="send" 
                    size="is-small"
                    v-model="ageArray"
                    native-value="age_range=under+10">
                        Under 10
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ageArray"
                    native-value="age_range=10-17">
                        10 to 17
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ageArray"
                    native-value="age_range=18-24">
                        18 to 24
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ageArray"
                    native-value="age_range=25-34">
                        24 - 34
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ageArray"
                    native-value="age_range=over+34">
                        Over 34
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ageArray"
                    native-value="age_range=Not+Stated">
                        Not Stated
                </b-checkbox>
            </div>
        </section>

        <p class="menu-label">Ethnicity</p>
        <section class="checkboxes">
            <div class="field">
                <b-checkbox 
                    @input="send"
                    size="is-small"
                    v-model="ethnicityArray"
                    native-value="ethnicity=Asian">
                        Asian
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ethnicityArray"
                    native-value="ethnicity=Black">
                        Black
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ethnicityArray"
                    native-value="ethnicity=Chinese+or+other">
                        Chinese or other
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ethnicityArray"
                    native-value="ethnicity=Mixed">
                        Mixed
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ethnicityArray"
                    native-value="ethnicity=Not+stated">
                        Not Stated
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="ethnicityArray"
                    native-value="ethnicity=White">
                        White
                </b-checkbox>
            </div>            
        </section>

    </div>
</template>

<script>
export default {
    data() {
        return {
            ageArray: ["age_range=under+10","age_range=10-17","age_range=18-24",
            "age_range=25-34","age_range=over+34","age_range=Not+Stated"],
            ethnicityArray: ["ethnicity=Chinese+or+other","ethnicity=Black","ethnicity=Not+stated","ethnicity=Mixed","ethnicity=White","ethnicity=Asian"]
        }
    },
    computed: {
        merge: function() {
            let arr = [...this.ageArray, ...this.ethnicityArray]
            return arr.join("&")
        }
    },
    methods: {
        send() {
            this.$emit('update-checkbox', this.merge)
        }
    },
    created: function () {
        this.$emit('update-checkbox', this.merge)
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
    max-height: 15vh;
    overflow-y: scroll;
    }
    .menu-label {
    margin: 1px 0
    }

</style>
