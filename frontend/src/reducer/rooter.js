import { combineReducers } from 'redux'
import passReducer from './reducer/pwReducer'
import userReducer from './reducer/usrReducer'


const rooter = combineReducers({
    user: userReducer,
    pass: passReducer,
})

export default rooter;

