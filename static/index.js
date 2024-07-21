htmx.on("form", "submit", (e) => {
	e.preventDefault();

	/** @type {FormData} */
	let values = htmx.values(htmx.find("form"));

	let params = new URLSearchParams(
		Object.fromEntries([...values.entries()]),
	).toString();

	htmx.ajax("GET", "/android", {
		target: "*[data-platform='android'] textarea",
		values: values,
	});
	htmx.ajax("GET", "/desktop", {
		target: "*[data-platform='desktop'] textarea",
		values: values,
	});

	htmx.find("*[data-platform='android'] a").setAttribute(
		"href",
		`/android?decimal=true&${params}`,
	);

	htmx.find("*[data-platform='android'] a").setAttribute(
		"download",
		`${values.get("name") || "my_theme"}.attheme`,
	);
});
