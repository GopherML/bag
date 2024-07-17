# Set up a 1200x600 terminal with 46px font.
Set FontSize 46
Set Width 1200
Set Height 600

Sleep 1s
Type "./bag-cli --training ../examples/yes-no-training.yaml --interactive"
Enter

Sleep 1s
Type "Yeah!"
Enter

Sleep 1s
Type "No way!"
Enter

Sleep 1s
Type "Bet"
Enter

Sleep 1s
Type "Nah fam"
Enter

# Admire the output for a bit.
Sleep 5s