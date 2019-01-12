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

        <p class="menu-label">Outcome</p>
        <section class="checkboxes">
            <div class="field">
                <b-checkbox 
                    @input="send"
                    size="is-small"
                    v-model="outcomeArray"
                    native-value="outcome=Arrest">
                        Arrest
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="outcomeArray"
                    native-value="outcome=Other">
                        Other
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="outcomeArray"
                    native-value="outcome=Community+resolution">
                        Community resolution
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="outcomeArray"
                    native-value="outcome=Nothing+Found">
                        Nothing Found
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="outcomeArray"
                    native-value="outcome=Caution">
                        Caution
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="outcomeArray"
                    native-value="outcome=Penalty+Notice">
                        Penalty Notice
                </b-checkbox>
            </div>           
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="outcomeArray"
                    native-value="outcome=Summons">
                        Summons
                </b-checkbox>
            </div>             
        </section>

        <p class="menu-label">Object of Search</p>
        <section class="checkboxes">
            <div class="field">
                <b-checkbox 
                    @input="send"
                    size="is-small"
                    v-model="objectOfSearchArray"
                    native-value="object_of_search=Criminal+damage">
                        Criminal damage
                </b-checkbox>
            </div>   
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="objectOfSearchArray"
                    native-value="object_of_search=Drugs">
                        Drugs
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="objectOfSearchArray"
                    native-value="object_of_search=Firearms">
                        Firearms
                </b-checkbox>
            </div>  
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="objectOfSearchArray"
                    native-value="object_of_search=Going+equipped">
                        Going equipped
                </b-checkbox>
            </div>      
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="objectOfSearchArray"
                    native-value="object_of_search=Offensive+weapons">
                        Offensive weapons
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="objectOfSearchArray"
                    native-value="object_of_search=Other">
                        Other
                </b-checkbox>
            </div>
            <div class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="objectOfSearchArray"
                    native-value="object_of_search=Stolen+property">
                        Stolen property
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
            ethnicityArray: ["ethnicity=Chinese+or+other","ethnicity=Black","ethnicity=Not+stated","ethnicity=Mixed","ethnicity=White","ethnicity=Asian"],
            objectOfSearchArray: ["object_of_search=Other","object_of_search=Offensive+weapons","object_of_search=Drugs","object_of_search=Going+equipped","object_of_search=Firearms","object_of_search=Criminal+damage","object_of_search=Stolen+property"],
            outcomeArray: ["outcome=Arrest","outcome=Other","outcome=Community+resolution","outcome=Nothing+Found","outcome=Caution","outcome=Penalty+Notice","outcome=Summons"]
        }
    },
    computed: {
        merge: function() {
            let arr = [...this.ageArray, ...this.ethnicityArray, ...this.outcomeArray, ...this.objectOfSearchArray]
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
    max-height: 125px;
    overflow-y: scroll;
    }
    .menu-label {
    margin: 1px 0
    }

</style>
