import { Box, MantineProvider, List, ThemeIcon } from "@mantine/core";
import {CheckCircleFillIcon} from "@primer/octicons-react";
import useSWR from "swr";
import './App.css';
import AddTodo from './components/AddTodo';

export interface Todo {
  id: number;
  title: string;
  body: string;
  done: boolean;
}
export const ENDPOINT = 'http://localhost:4000'

const fetcher = (url: string) => fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

// const fetcher = function(url) {
//   return fetch(`${ENDPOINT}/${url}`).then(function(r) {
//     return r.json();
//   });
// };

function App() {
  //const {data, mutate} is destructuring assignment
  const {data, mutate} = useSWR<Todo[]>('api/todos', fetcher)

  

  async function markTodoAsDone(id: number){
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/done`, {
      method: 'PATCH'
    }).then((r) => r.json());
    mutate(updated);
  }

  console.log(data);

  return (
    <MantineProvider>
      <Box
      sx={(theme) => ({
        padding: "2rem",
        width:'100%',
        maxWidth: '40rem',
        margin: "0 auto",
      })}
      >
        <List spacing="xs" size="sm" mb={12} center>
        {data?.map((todo) => {
          return <List.Item 
          onClick={() => markTodoAsDone(todo.id)}
          key={`todo_list__${todo.id}`}
          icon={
            todo.done ? (
              <ThemeIcon color="teal" size={24} radius="xl">
                <CheckCircleFillIcon size={20}/>
              </ThemeIcon>
            ) : (
              <ThemeIcon color="gray" size={24} radius="xl">
                <CheckCircleFillIcon size={20}/>
              </ThemeIcon>
            )
          }
          >{todo.title}</List.Item>
        })} 
        </List>

<AddTodo mutate={mutate} />
      </Box>

      
    </MantineProvider>
  );
}

export default App;