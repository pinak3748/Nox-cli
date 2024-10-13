package content

import (
	"strings"
)

func GenerateMainPageContent(name string) string {
	template := `import React from 'react';

interface {pageName}Props {
  // Add any props here
}

const {pageName}: React.FC<{pageName}Props> = () => {
  return (
    <div className={styles.container}>
      <h1>{pageName}</h1>
      <p>Welcome to the {pageName} page!</p>
      {/* Add your page content here */}
    </div>
  );
};

export default {pageName};
`

	return strings.ReplaceAll(template, "{pageName}", name)
}
