import thunk from 'redux-thunk'
import { applyMiddleware, createStore } from 'redux'
import rooter from './rooter'



const store = createStore(rooter, applyMiddleware(thunk))

export default store