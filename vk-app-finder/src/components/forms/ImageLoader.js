import React, {PureComponent} from 'react';
import ReactCrop from 'react-image-crop';
import Icon24Camera from '@vkontakte/icons/dist/24/camera';
import 'react-image-crop/dist/ReactCrop.css';
import {Div, File, FormLayoutGroup} from "@vkontakte/vkui";
import FormStatus from "@vkontakte/vkui/dist/components/FormStatus/FormStatus";

class ImageLoader extends PureComponent {
  state = {
    src: this.props.image,
    crop: this.props.crop,
  };

  onSelectFile = e => {
    if (e.target.files && e.target.files.length > 0) {
      const reader = new FileReader();
      reader.addEventListener('load', () => {
        this.setState({src: reader.result});
        this.props.onImageChange(reader.result);
      });
      reader.readAsDataURL(e.target.files[0]);
    }
  };

  // If you setState the crop in here you should return false.
  onImageLoaded = image => {
    this.imageRef = image;
  };

  onCropComplete = crop => {
    this.makeClientCrop(crop);
  };

  onCropChange = (crop, percentCrop) => {
    this.props.onCropChange(crop);
    // You could also use percentCrop:
    // this.setState({ crop: percentCrop });
    this.setState({crop});
  };

  async makeClientCrop(crop) {
    if (this.imageRef && crop.width && crop.height) {
      const croppedImageUrl = await this.getCroppedImg(
        this.imageRef,
        crop,
        'newFile.jpeg'
      );
      this.setState({croppedImageUrl});
    }
  }

  getCroppedImg(image, crop, fileName) {
    const canvas = document.createElement('canvas');
    const scaleX = image.naturalWidth / image.width;
    const scaleY = image.naturalHeight / image.height;
    canvas.width = crop.width;
    canvas.height = crop.height;
    const ctx = canvas.getContext('2d');

    ctx.drawImage(
      image,
      crop.x * scaleX,
      crop.y * scaleY,
      crop.width * scaleX,
      crop.height * scaleY,
      0,
      0,
      crop.width,
      crop.height
    );

    return new Promise((resolve, reject) => {
      canvas.toBlob(blob => {
        if (!blob) {
          //reject(new Error('Canvas is empty'));
          console.error('Canvas is empty');
          return;
        }
        blob.name = fileName;
        window.URL.revokeObjectURL(this.fileUrl);
        this.fileUrl = window.URL.createObjectURL(blob);
        this.props.onPictureSet(blob);
        resolve(this.fileUrl);
      }, 'image/jpeg');
    });
  }

  render() {
    const {croppedImageUrl} = this.state;
    const imageWrapperStyles = {
      marginLeft: 'auto',
      marginRight: 'auto',
      display: 'block',
      width: 'max-content'
    };

    return (
      <FormLayoutGroup>
        <FormStatus header="Фотография животного">
          Фотография поможет другим узнать животное, а мы попробуем подсказать его породу.
        </FormStatus>
        {this.props.image && <Div>
          <ReactCrop src={this.props.image}
                     crop={this.props.crop}
                     ruleOfThirds
                     keepSelection
                     style={imageWrapperStyles}
                     imageStyle={{maxHeight: '345px', width: 'auto'}}
                     onImageLoaded={this.onImageLoaded}
                     onComplete={this.onCropComplete}
                     onChange={this.onCropChange}/>
          <div style={{marginTop: '15px', width: '100%', textAlign: 'center'}}
               className={'form__label'}>
            {'Выберите область изображения, на которой лучше всего видно найденного питомца'}
          </div>
        </Div>}
        <File accept="image/*"
              controlSize={'xl'}
              before={<Icon24Camera/>}
              onChange={this.onSelectFile}>
          Выберите изображение
        </File>

        {/*{false && (*/}
          {/*<img alt="Crop" style={{maxWidth: '100%'}} src={croppedImageUrl}/>*/}
        {/*)}*/}
      </FormLayoutGroup>
    );
  }
}

export default ImageLoader;