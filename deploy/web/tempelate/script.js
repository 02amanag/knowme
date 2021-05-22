$.fn.id_changer = function(new_id) {
    return this.each(function() {
        $(this).find('input').each(function() {
            var ob = $(this);
            ob.attr('id', this.id.replace(/_\d$/, '_' + new_id));
            ob.attr('name', this.name.replace(/_\d$/, '_' + new_id));
        });
    });
};


// Add a new repeating section
$('.addFlight').click(function() {
    var lastRepeatingGroup = $('.repeatingSection').last();
    var cloned = lastRepeatingGroup.clone()

    //TODO change 3 to a variable later...
    cloned = cloned.id_changer(3);

    cloned.insertAfter(lastRepeatingGroup);
    return false;
});


// Delete a repeating section
$('.deleteFight').click(function() {
    $(this).parent('div').remove();
    return false;
});