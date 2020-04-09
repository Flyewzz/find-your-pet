import {ModalPageHeader, PanelHeaderButton} from "@vkontakte/vkui";
import Icon24Cancel from "@vkontakte/icons/dist/24/cancel";
import Icon24Done from "@vkontakte/icons/dist/24/done";
import React from "react";

const FilterHeader = (props) => {
  return (
    <ModalPageHeader
      left={
        <PanelHeaderButton onClick={props.onClose}>
          <Icon24Cancel/>
        </PanelHeaderButton>
      }
      right={<PanelHeaderButton onClick={props.onDone}>
        <Icon24Done/>
      </PanelHeaderButton>}
    >
      Фильтры
    </ModalPageHeader>
  );
};

export default FilterHeader;