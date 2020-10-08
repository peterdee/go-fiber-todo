/**
 * Handle TODO editing: open a modal
 * @param {*} event - click event
 * @returns {Promise<void>}
 */
const handleEdit = async (event) => {
  // clear the error message
  $('#todo-error').empty();

  // get item ID
  const { id: fullId = '' } = event.target;
  if (!fullId) {
    return $('#todo-error').append('Error editing an item!');
  }
  const [id = ''] = fullId.split('-').slice(-1);
  if (!id) {
    return $('#todo-error').append('Error editing an item!');
  }

  // create modal window
  $('#modals').empty().append(`
<div class="background"></div>
<div class="flex direction-column modal-window">
  <div class="fs-20 mb-16 text-center noselect">
    Edit TODO
  </div>
  <form id="edit-form">
    <textarea
      class="edit-textarea mb-16"
      id="edit-textarea"
    ></textarea>
    <div
      class="todo-error text-center noselect"
      id="edit-error"
    ></div>
    <div class="flex justify-content-space-between noselect">
      <button
        class="edit-button success"
        id="edit-save"
        type="submit"
      >
        SAVE
      </button>
      <button
        class="edit-button fail"
        id="edit-cancel"
        type="button"
      >
        CANCEL
      </button>
    </div>
  </form>
</div>
  `);

  // backdrop for the modal
  $('.background').on('click', () => $('#modals').empty());

  // close modal on 'Cancel' click
  $('#edit-cancel').on('click', () => $('#modals').empty());

  try {
    // load an item when opening a modal
    const { data: { completed = false, text = '' } = {} } = await $.ajax({
      method: 'GET',
      url: `/api/todos/get/${id}`,
    });

    // put the text in the textarea
    $('#edit-textarea').val(text);

    // handle saving
    $('#edit-form').on('submit', async (event) => {
      event.preventDefault();

      const newText = $('#edit-textarea').val();
      if (!(newText && newText.trim())) {
        return;
      }

      try {
        await $.ajax({
          data: {
            completed,
            text: newText,
          },
          method: 'PATCH',
          url: `/api/todos/update/${id}`,
        });
  
        // reload the page
        return window.location.reload();
      } catch {
        // show the error and disable the form
        $('#edit-error').empty().append('Error updating the data!');
        $('#edit-save').attr('disabled', true);
        $('#edit-form').attr('disabled', true);
      }
    });
  } catch {
    // show the error and disable the form
    $('#edit-error').empty().append('Error loading the data!');
    $('#edit-save').attr('disabled', true);
    $('#edit-form').attr('disabled', true);
  }
};

/**
 * Handle Todo item click (change the 'completed' status)
 * @param {*} event - click event
 * @returns {Promise<void>}
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
  $('.edit-todo').on('click', handleEdit);
  $('#todo-form').on('submit', handleSubmit);
});
