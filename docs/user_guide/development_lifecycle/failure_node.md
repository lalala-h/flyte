(failure_node)=
# Failure node

```{eval-rst}
 .. tags:: FailureNode, Intermediate
```

:::{warning}
This feature is only available starting in Flyte 1.10.7.
:::


The failure node feature enables you to designate a specific node to execute in the event of a failure within your workflow.

For example, a workflow involves creating a cluster at the beginning, followed by the execution of tasks, and concludes with the deletion of the cluster once all tasks are completed. However, if any task within the workflow encounters an error, flyte will abort the entire workflow and won’t delete the cluster. This poses a challenge if you still need to clean up the cluster even in a task failure.

To address this issue, you can add a failure node into your workflow. This ensures that critical actions, such as deleting the cluster, are executed even in the event of failures occurring throughout the workflow execution

```{note}
To clone and run the example code on this page, see the [Flytesnacks repo][flytesnacks].
```

```{literalinclude} /examples/development_lifecycle/development_lifecycle/failure_node.py
:caption: development_lifecycle/failure_node.py
:lines: 1-7
```

Create a task that will fail during execution:

```{literalinclude} /examples/development_lifecycle/development_lifecycle/failure_node.py
:caption: development_lifecycle/failure_node.py
:lines: 11-19
```

Create a task that will be executed if any of the tasks in the workflow fail:

```{literalinclude} /examples/development_lifecycle/development_lifecycle/failure_node.py
:caption: development_lifecycle/failure_node.py
:pyobject: clean_up
```

Specify the `on_failure` to a cleanup task. This task will be executed if any of the tasks in the workflow fail:

:::{note}
The inputs of `clean_up` must exactly match the workflow's inputs. Additionally, the `err` parameter will be 
populated with the error message encountered during execution.
:::

```{literalinclude} /examples/development_lifecycle/development_lifecycle/failure_node.py
:caption: development_lifecycle/failure_node.py
:pyobject: subwf
```

By setting the failure policy to `FAIL_AFTER_EXECUTABLE_NODES_COMPLETE` to ensure that the `wf1` is executed even if the subworkflow fails. In this case, both parent and child workflows will fail, resulting in the `clean_up` task being executed twice:

```{literalinclude} /examples/development_lifecycle/development_lifecycle/failure_node.py
:caption: development_lifecycle/failure_node.py
:lines: 43-54
```

You can also set the `on_failure` to a workflow. This workflow will be executed if any of the tasks in the workflow fail:

```{literalinclude} /examples/development_lifecycle/development_lifecycle/failure_node.py
:caption: development_lifecycle/failure_node.py
:pyobject: wf2
```

[flytesnacks]: https://github.com/flyteorg/flytesnacks/tree/master/examples/development_lifecycle/
