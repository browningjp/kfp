# `kfp`: Kubeflow Pipelines SDK

[![PyPI Package version](https://badge.fury.io/py/kfp.svg)](https://badge.fury.io/py/kfp)
[![PyPI Python Version](https://img.shields.io/pypi/pyversions/kfp.svg)](https://pypi.org/project/kfp/)
[![PyPI Downloads](https://img.shields.io/pypi/dm/kfp)](https://pypi.org/project/kfp/)
[![Documentation Status](https://readthedocs.org/projects/kubeflow-pipelines/badge/?version=latest)](https://kubeflow-pipelines.readthedocs.io/en/stable/?badge=latest)
[![Code Style](https://img.shields.io/badge/code%20style-yapf-brightgreen.svg)](https://github.com/google/yapf)

Kubeflow Pipelines is a platform for building and deploying portable, scalable machine learning workflows based on Docker containers within the [Kubeflow](https://www.kubeflow.org/) project.

Use Kubeflow Pipelines to compose a multi-step workflow ([pipeline](https://www.kubeflow.org/docs/components/pipelines/concepts/pipeline/)) as a [graph](https://www.kubeflow.org/docs/components/pipelines/concepts/graph/) of containerized [tasks](https://www.kubeflow.org/docs/components/pipelines/concepts/step/) using [Python code](https://www.kubeflow.org/docs/components/pipelines/sdk/python-function-components/#getting-started-with-python-function-based-components) and/or [YAML](https://www.kubeflow.org/docs/components/pipelines/sdk/component-development/#creating-a-component-specification). Then, [run](https://www.kubeflow.org/docs/components/pipelines/concepts/run/) your pipeline with specified pipeline arguments, rerun your pipeline with new arguments or data, [schedule](https://www.kubeflow.org/docs/components/pipelines/concepts/run-trigger/) your pipeline to run on a recurring basis, organize your runs into [experiments](https://www.kubeflow.org/docs/components/pipelines/concepts/experiment/), save machine learning artifacts to compliant [artifact registries](https://www.kubeflow.org/docs/components/pipelines/concepts/metadata/), and visualize it all through the [Kubeflow Dashboard](https://www.kubeflow.org/docs/components/central-dash/overview/).

## Documentation
* [Kubeflow Pipelines Overview](https://www.kubeflow.org/docs/components/pipelines/introduction/)
* [SDK Overview](https://www.kubeflow.org/docs/components/pipelines/sdk/sdk-overview/)
* [SDK API Documentation](https://kubeflow-pipelines.readthedocs.io/en/stable/)

## Installation

To install the latest stable release, run:

```sh
pip install kfp
```

## Getting started

The following is an example of a simple pipeline with one Python function-based component used in two separate tasks to do basic addition:

```python
import kfp
from kfp.components import create_component_from_func
import kfp.dsl as dsl

def add(a: float, b: float) -> float:
    '''Calculates sum of two arguments'''
    return a + b


# create a component using the add function
add_op = create_component_from_func(add)

# compose your pipeline using the dsl.pipeline decorator
@dsl.pipeline(
    name='Addition pipeline',
    description='An example pipeline that performs addition calculations.')
def add_pipeline(
    a='1',
    b='7',
):
    first_add_task = add_op(a, 4)
    second_add_task = add_op(first_add_task.output, b)

# instantiate a client and submit your pipeline with arguments
client = kfp.Client(host='<my-host-url>')
client.create_run_from_pipeline_func(
    add_pipeline, arguments={
        'a': '7',
        'b': '8'
    })

```

For more information, refer to [Building Python function-based components](https://www.kubeflow.org/docs/components/pipelines/sdk/python-function-components/).