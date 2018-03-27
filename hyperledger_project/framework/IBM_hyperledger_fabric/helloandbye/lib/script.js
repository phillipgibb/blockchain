/**
 * Hello and Bye
 * 
 */

/**
@param {org.acme.vocab.Hello} word
@transaction
*/

function hello(word) {
  console.log("Message: " + word.test.message);
}

/**
@param {org.acme.vocab.Bye} word
@transaction
*/


function bye(word) {
    console.log("Message: " + word.test_t.message)
}