import axios from 'axios'; // Importa axios para realizar peticiones HTTP

const instance = axios.create({ // Crea una instancia de axios
    baseURL: "http://localhost:3000/items", // URL base para las peticiones
});

export default instance; // Crea una instancia de axios con la URL base
export const api = instance; //Exporta la instancia para uso en otros archivos