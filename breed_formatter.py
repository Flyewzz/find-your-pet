

def breed_format(breed_path):
    breed_with_number = breed_path.split('/')[-1]
    # Eng breed
    return breed_with_number.split('.')[-1]