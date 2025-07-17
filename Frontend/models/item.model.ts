export interface Item {
    ticker: string; // Ticker del ítem
    target_from: string; // Valor objetivo desde
    target_to: string; // Valor objetivo hasta
    company: string; // Nombre de la empresa
    action: string; // Acción a realizar (compra, venta, etc.)
    brokerage: string; // Nombre del bróker
    rating_from: string; // Calificación desde
    rating_to: string; // Calificación hasta
    time: string; // Tiempo de la acción
    pageCount: number; // Contador de páginas
}