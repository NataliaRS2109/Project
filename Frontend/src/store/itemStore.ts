// stores/itemStore.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Item } from '../models/item.model'
import { api } from '../utils/axios'

export const useItemStore = defineStore('item', () => {
  const page = ref(1)
  const items = ref<Item[]>([])
  const itemsRecommendation = ref<Item[]>([])
  const totalItems = ref(0)
  const isLoading = ref(false)

  const fetchItems = async () => {
    try {
      isLoading.value = true
      const { data } = await api.get(`/items?page=${page.value}`)
      const { data: recommendedData } = await api.get('/items/recommended')
      items.value = data
      itemsRecommendation.value = recommendedData
      totalItems.value = data.length
    } catch (error) {
      console.error('Error fetching items:', error)
    } finally {
      isLoading.value = false
    }
  }

  const changePage = (newPage: number) => {
    page.value = newPage
    fetchItems()
  }

  return {
    page,
    items,
    itemsRecommendation,
    totalItems,
    isLoading,
    fetchItems,
    changePage
  }
})
