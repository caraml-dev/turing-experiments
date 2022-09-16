import React from "react";

import { EuiFlexItem, EuiText, EuiFormRow, EuiFieldText } from "@elastic/eui";
import { FormLabelWithToolTip } from "@gojek/mlp-ui";

import { Panel } from "components/panel/Panel";

export const RouteNamePathPanel = ({
  routeNamePath,
  onChange,
  errors,
}) => {
  return (
    <Panel title={"Route Name Path"}>
      <EuiFlexItem>
        <EuiFormRow
          fullWidth
          label={
            <FormLabelWithToolTip
              label="Route Name Path *"
              content="Specify the path in the treatment configuration where the route name for the final response can be found."
            />
          }
          isInvalid={!!errors}
          error={errors}
          display="row">
          <EuiFieldText
            fullWidth
            placeholder="policy.route_name"
            value={routeNamePath}
            onChange={(e) => onChange(e.target.value)}
            isInvalid={!!errors}
            name="route-name-path"
            prepend={<EuiText size={"s"}>treatment.configuration.</EuiText>}
          />
        </EuiFormRow>
      </EuiFlexItem>
    </Panel>
  );
};
