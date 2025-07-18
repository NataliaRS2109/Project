<script setup lang="ts">
    import type { Item } from '../models/item.model'; // Import the Item type from the model
    // import { defineProps } from 'vue'; // Import defineProps to define component properties

    // Define the component properties using defineProps
    // Define an interface Props to type the component properties.
    interface Props {
        item: Item;
    }

    defineProps<Props>(); // Define the component properties using the Props interface. This allows the component to receive an array of Item objects as a property.
</script>


<template>
    <div class="h-fit text-left flex flex-col justify-around p-4">
        <h4 class="text-2xl font-bold">{{ item.ticker }} - {{ item.company }}</h4>
        <p><span class="font-bold">Brokerage:</span> {{ item.brokerage }}</p>
        <button class="border border-purple-800 rounded px-2 py-1 mt-4 w-20 text-white bg-purple-900 hover:bg-purple-700">Details</button>
        <!-- <p><span class="font-bold">Action:</span> {{ item.action }}</p>
        <p><span class="font-bold">Rating From:</span> {{ item.rating_from }} - <span class="font-bold">Rating To:</span> {{ item.rating_to }}</p>
        <p><span class="font-bold">Target From:</span> {{ item.target_from }} - <span class="font-bold">Target To:</span> {{ item.target_to }}</p> -->
    </div>
</template>