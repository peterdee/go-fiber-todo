const handleComplete = async (event) => {

};

/**
 * Delete a Todo item
 * @param {*} event - button click event
 * @returns {Promise<void>}
 */
const handleDelete = async (event) => {
  // clear the error message
  $('#todo-error').empty();

  // get item ID
  const { id: fullId = '' } = event.target;
  if (!fullId) {
    return $('#todo-error').append('Error deleting an item!');
  }
  const [id = ''] = fullId.split('-').slice(-1);
  if (!id) {
    return $('#todo-error').append('Error deleting an item!');
  }

  try {
    // send a request
    await $.ajax({
      method: 'DELETE',
      url: `/api/todos/delete/${id}`,
    });

    // remove an element from the DOM
    return $(`#todo-${id}`).remove();
  } catch (error) {
    return $('#todo-error').append('Error deleting an item!');
  }
};

/**
 * Handle form submitting
 * @param {*} event - form submit event
 * @returns {Promise<void>}
 */
const handleSubmit = async (event) => {
  event.preventDefault();
  
  // clear the error message
  $('#todo-error').empty();

  // check the input
  const text = $('#todo-input').val();
  $('#todo-input').val('');
  if (!(text && text.trim())) {
    return $('#todo-error').append('Please provide the text!');
  }

  // send a request
  try {
    const { data } = await $.ajax({
      data: {
        text,
      },
      method: 'POST',
      url: '/api/todos/add',
    });

    // add an element to the list
    return $('#todos-list').append(`
<div
  id="todo-${data.id}"
  class="flex justify-content-space-between mb-16"
>
  <div>
    ${data.text}
  </div>
  <div>
    Completed: ${data.completed}
  </div>
  <button
    class="delete-todo"
    id="todo-delete-${data.id}"
    type="button"
  >
    Delete
  </button>
</div>
    `);
  } catch (error) {
    return $('#todo-error').append('Error adding an item!');
  }
};

/**
 * On page load
 */
$(document).ready(() => {
  $('.complete-todo').on('submit', handleComplete);
  $('.delete-todo').on('click', handleDelete);
  $('#todo-form').on('submit', handleSubmit);
});
