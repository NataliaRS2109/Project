<script setup lang="ts">
    import type { Item } from '../models/item.model'; // Importa el tipo Item desde el modelo
    import { defineProps } from 'vue'; // Importa defineProps para definir las propiedades del componente
    // Define las propiedades del componente usando defineProps

    // Define una interfaz Props para tipar las propiedades del componente.
    interface Props { // Define una interfaz Props para tipar las propiedades del componente.
        items: Item[];
    }

    defineProps<Props>(); // Define las propiedades del componente usando la interfaz Props. Esto permite que el componente reciba un array de objetos Country como propiedad.
</script>

<template>
    <div class="relative flex flex-col w-full h-full overflow-auto bg-white shadow-md rounded-lg bg-clip-border mt-4">
        <table class="w-full table-auto min-w-max text-left">
            <thead>
                <tr>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300">
                        <p class="text-sm font-normal leading-none text-slate-800">
                        TICKER
                        </p>
                    </th>            
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300">
                        <p class="text-sm font-normal leading-none text-slate-800">
                        COMPANY
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300">
                        <p class="text-sm font-normal leading-none text-slate-800">
                        BROKERAGE
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300">
                        <p class="text-sm font-normal leading-none text-slate-800">
                        ACTION
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300">
                        <p class="text-sm font-normal leading-none text-slate-800">
                        RATING_FROM
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300">
                        <p class="text-sm font-normal leading-none text-slate-800">
                        RATING_TO
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300">
                        <p class="text-sm font-normal leading-none text-slate-800">
                        TARGET_FROM
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300">
                        <p class="text-sm font-normal leading-none text-slate-800">
                        TARGET_TO
                        </p>
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in items" class="hover:bg-gray-100">
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.ticker }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.company }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.brokerage }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.action }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.rating_from }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.rating_to }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.target_from }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.target_to }}
                        </p>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>


