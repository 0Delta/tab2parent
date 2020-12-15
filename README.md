tab2parent
========

Usage:
------
testfile
```txt
metadata
	annotations
	managedFields
		apiVersion
		fieldsType
	spec
		replicas
		selector
			matchExpressions
				key
				operator
				values
			matchLabels
		strategy
```

```bash
$ tab2parent testfile
metadata
metadata.annotations
metadata.managedFields
metadata.managedFields.apiVersion
metadata.managedFields.fieldsType
metadata.spec
metadata.spec.replicas
metadata.spec.selector
metadata.spec.selector.matchExpressions
metadata.spec.selector.matchExpressions.key
metadata.spec.selector.matchExpressions.operator
metadata.spec.selector.matchExpressions.values
metadata.spec.selector.matchLabels
metadata.spec.strategy
```

Requirements:
-------------
+ go

Install:
--------
+ `go install github.com/0Delta/tab2parent`

License:
--------
MIT

Author:
-------
0Δ(@0Delta)
