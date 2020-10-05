/**
 * Handle Todo item click (change the 'completed' status)
 * @param {*} event - click event
 */
const handleComplete = async (event) => {
  const { id = '' } = event.target;
  if (!id) {
    return $('#todo-error').append('Error deleting an item!');
  }

  const completed = $(`#${id}`).attr('class').split(' ').includes('completed');

  try {
    // send a request
    await $.ajax({
      data: {
        completed: !completed,
      },
      method: 'PATCH',
      url: `/api/todos/update/${id}`,
    });

    // reload the page
    return window.location.reload();
  } catch {
    return $('#todo-error').append('Error deleting an item!');
  }
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
  } catch {
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
    await $.ajax({
      data: {
        text,
      },
      method: 'POST',
      url: '/api/todos/add',
    });

    // reload the page
    return window.location.reload();
  } catch (error) {
    return $('#todo-error').append('Error adding an item!');
  }
};

/**
 * On page load
 */
$(document).ready(() => {
  $('.complete-todo').on('click', handleComplete);
  $('.delete-todo').on('click', handleDelete);
  $('#todo-form').on('submit', handleSubmit);
});
