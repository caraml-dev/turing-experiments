import "experiments/list/search/SearchExperimentsFilters.scss";

import React, { useContext, useEffect, useState } from "react";

import {
  EuiButtonEmpty,
  EuiFlexGroup,
  EuiFlexItem,
  EuiFlyoutBody,
  EuiFlyoutHeader,
  EuiForm,
  EuiSpacer,
  EuiTitle,
} from "@elastic/eui";
import isEqual from "lodash/isEqual";

import {
  experimentStatuses,
  experimentTiers,
  experimentTypes,
} from "experiments/components/typeOptions";
import ExperimentDateFilter from "experiments/list/search/components/ExperimentDateFilter";
import ExperimentSegmenterFilter from "experiments/list/search/components/ExperimentSegmenterFilter";
import ExperimentSegmenterMatchOptionsFilter from "experiments/list/search/components/ExperimentSegmenterMatchOptionsFilter";
import ExperimentTypeOptionsFilter from "experiments/list/search/components/ExperimentTypeOptionsFilter";
import ExperimentUpdatedByFilter from "experiments/list/search/components/ExperimentUpdatedByFilter";
import ExperimentSearchContext from "experiments/list/search/context";
import schema from "experiments/list/search/validation/schema";
import SegmenterContext from "providers/segmenters/context";
import { extractErrors } from "utils/helpers";

const SearchExperimentFilters = ({ onChange }) => {
  const { clearFilters, getFilter, getFilters, isFilterSet, setFilter } =
    useContext(ExperimentSearchContext);
  const appliedFilters = getFilters();
  const { dependencyMap, getSegmenterOptions } = useContext(SegmenterContext);
  const segmenterOptions = getSegmenterOptions(appliedFilters).filter(
    (opt) => opt.options.length !== 0
  );

  const [validationErrors, setValidationState] = useState({
    filters: {},
    errors: {},
  });
  useEffect(() => {
    if (!isEqual(validationErrors.filters, appliedFilters)) {
      // Validate the filters
      schema
        .validate(appliedFilters, { abortEarly: false })
        .catch(function (err) {
          if (err.errors) {
            // Update filters and errors in the validation state
            setValidationState({
              filters: appliedFilters,
              errors: extractErrors(err),
            });
          }
        });
      // No errors, update filters and clear errors
      setValidationState({ filters: appliedFilters });
    }
  }, [appliedFilters, validationErrors]);

  const onChangeHandler = (key) => (value) => {
    onChange();
    const dependentFiltersToDelete = dependencyMap[key]
      ? dependencyMap[key]
      : [];
    setFilter(key, value, dependentFiltersToDelete);
  };

  // Clear all filters
  const clearAllFilters = () => {
    clearFilters();
  };

  return (
    <>
      <EuiFlyoutHeader hasBorder className="searchPanelFlyoutHeader">
        <EuiFlexGroup alignItems="center">
          <EuiFlexItem>
            <EuiTitle size="s">
              <h4>Filters</h4>
            </EuiTitle>
          </EuiFlexItem>
          <EuiFlexItem>
            <EuiButtonEmpty
              iconSide="right"
              onClick={() => clearAllFilters()}
              iconType="trash"
              isDisabled={!isFilterSet()}
              size="s">
              Reset
            </EuiButtonEmpty>
          </EuiFlexItem>
        </EuiFlexGroup>
      </EuiFlyoutHeader>
      <EuiFlyoutBody className="searchPanelFlyoutBody">
        <EuiFlexGroup direction="row" justifyContent="center">
          <EuiFlexItem grow={false} style={{ width: "75%" }}>
            <EuiForm>
              <ExperimentTypeOptionsFilter
                label="Experiment Type"
                options={experimentTypes}
                value={getFilter("type")}
                onChange={onChangeHandler("type")}
              />
            </EuiForm>
            <EuiSpacer size="m" />
            <EuiForm>
              <ExperimentTypeOptionsFilter
                label="Experiment Tier"
                options={experimentTiers}
                value={getFilter("tier")}
                onChange={onChangeHandler("tier")}
              />
            </EuiForm>
            <EuiSpacer size="m" />
            <EuiForm>
              <ExperimentTypeOptionsFilter
                label="Experiment Status"
                options={experimentStatuses}
                value={getFilter("status")}
                onChange={onChangeHandler("status")}
              />
            </EuiForm>
            <EuiSpacer size="m" />
            <EuiForm>
              <ExperimentDateFilter
                startTime={getFilter("start_time")}
                endTime={getFilter("end_time")}
                onChange={onChangeHandler}
                errors={validationErrors.errors}
              />
            </EuiForm>
            <EuiSpacer size="m" />
            <EuiForm>
              <ExperimentUpdatedByFilter
                value={getFilter("updated_by")}
                onChange={onChangeHandler("updated_by")}
              />
            </EuiForm>
            <EuiSpacer size="m" />
            <EuiForm>
              {segmenterOptions.map((opt) => (
                <ExperimentSegmenterFilter
                  key={opt.name}
                  name={opt.name}
                  filteredOptions={opt.options}
                  onChange={onChangeHandler(opt.name)}
                  isMultiValued={opt.multi_valued}
                  value={getFilter(opt.name)}
                />
              ))}
              <ExperimentSegmenterMatchOptionsFilter
                value={getFilter("include_weak_match")}
                onChange={onChangeHandler("include_weak_match")}
              />
            </EuiForm>
            <EuiSpacer size="m" />
          </EuiFlexItem>
        </EuiFlexGroup>
      </EuiFlyoutBody>
    </>
  );
};

export default SearchExperimentFilters;
